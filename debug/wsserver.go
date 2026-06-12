// Package debug provides a WebSocket server that streams behavior tree
// execution to external tools, such as the Behavior Tree Editor's real-time
// debugger.
//
// WSServer implements core.Debugger. Attach it to a tree and tick as usual:
//
//	dbg := debug.NewWSServer(":6112")
//	defer dbg.Close()
//	tree.SetDebug(dbg)
//	for { tree.Tick(target, board) }
//
// Clients connect to ws://host:port/debug. The wire protocol is documented in
// behavior-tree-editor/docs/REALTIME_DEBUGGING.md, which is the single source
// of truth shared by both repos.
package debug

import (
	"encoding/json"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	bt "github.com/henrytien/behavior-tree"
)

// defaultFlushInterval coalesces status changes to ~10 Hz. A 60 fps game would
// otherwise flood the channel and the editor's renderer.
const defaultFlushInterval = 100 * time.Millisecond

// sendBuffer bounds each client's outbound queue. If a slow client falls
// behind, frames are dropped rather than blocking the tick or flusher — debug
// data is disposable and the next frame carries fresh state.
const sendBuffer = 16

// helloMsg is sent once to each client on connect so it can verify it is
// viewing the same tree it has open.
type helloMsg struct {
	Type      string `json:"type"`
	TreeID    string `json:"treeId"`
	NodeCount int    `json:"nodeCount"`
}

// tickMsg carries the node statuses that changed since the last flush, keyed by
// the node id the editor exported.
type tickMsg struct {
	Type   string            `json:"type"`
	TreeID string            `json:"treeId"`
	Seq    uint64            `json:"seq"`
	Nodes  map[string]string `json:"nodes"`
}

// blackboardMsg carries a snapshot of runtime variables for the variable panel.
type blackboardMsg struct {
	Type   string                 `json:"type"`
	TreeID string                 `json:"treeId"`
	Data   map[string]interface{} `json:"data"`
}

// client is one connected editor. A dedicated writer goroutine drains send, so
// only one goroutine ever writes to conn (gorilla requires serialized writes).
type client struct {
	conn *websocket.Conn
	send chan []byte
}

// WSServer streams tick events to connected clients. The zero value is not
// usable; construct it with NewWSServer.
type WSServer struct {
	flushInterval time.Duration
	upgrader      websocket.Upgrader
	httpServer    *http.Server

	mu      sync.Mutex
	clients map[*client]struct{}
	treeID  string
	// last holds the most recently reported status string per node, used to
	// detect changes. pending holds changes not yet flushed.
	last    map[string]string
	pending map[string]string
	seq     uint64

	done   chan struct{}
	closed atomic.Bool
}

// NewWSServer creates a WSServer and starts listening on addr (e.g. ":6112").
// Clients connect to the /debug path. Call Close to stop.
func NewWSServer(addr string) *WSServer {
	s := newServer(defaultFlushInterval)

	mux := http.NewServeMux()
	mux.HandleFunc("/debug", s.handleWS)
	s.httpServer = &http.Server{Addr: addr, Handler: mux}

	go func() { _ = s.httpServer.ListenAndServe() }()
	go s.flushLoop()
	return s
}

// newServer builds a server without binding a port. Used by NewWSServer and by
// tests (which drive it through httptest), so the flush interval is injectable.
func newServer(flushInterval time.Duration) *WSServer {
	return &WSServer{
		flushInterval: flushInterval,
		upgrader: websocket.Upgrader{
			// Local debugging tool: accept any origin. Bind to localhost in
			// production deployments if network exposure is a concern.
			CheckOrigin: func(*http.Request) bool { return true },
		},
		clients: make(map[*client]struct{}),
		last:    make(map[string]string),
		pending: make(map[string]string),
		done:    make(chan struct{}),
	}
}

// statusString maps a runtime status to the protocol enum.
func statusString(s bt.Status) string {
	switch s {
	case bt.SUCCESS:
		return "success"
	case bt.FAILURE:
		return "failure"
	case bt.RUNNING:
		return "running"
	case bt.ERROR:
		return "error"
	default:
		return "unknown"
	}
}

// --- core.Debugger ---

// OnTickStart records the tree being ticked. The first tree id seen is the one
// reported to clients.
func (s *WSServer) OnTickStart(treeID string) {
	s.mu.Lock()
	s.treeID = treeID
	s.mu.Unlock()
}

// OnNodeStatus records a node's status, queueing it for the next flush only if
// it changed since the last reported value.
func (s *WSServer) OnNodeStatus(treeID, nodeID string, status bt.Status) {
	str := statusString(status)
	s.mu.Lock()
	if s.last[nodeID] != str {
		s.last[nodeID] = str
		s.pending[nodeID] = str
	}
	s.mu.Unlock()
}

// OnTickEnd is a no-op: the flush loop owns send timing so a high tick rate
// does not translate into a high send rate.
func (s *WSServer) OnTickEnd(treeID string) {}

// --- flushing ---

// flushLoop sends coalesced status changes at flushInterval until Close.
func (s *WSServer) flushLoop() {
	ticker := time.NewTicker(s.flushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-s.done:
			return
		case <-ticker.C:
			s.flush()
		}
	}
}

// flush broadcasts pending changes, if any, as a single tick message and bumps
// the sequence number.
func (s *WSServer) flush() {
	s.mu.Lock()
	if len(s.pending) == 0 {
		s.mu.Unlock()
		return
	}
	s.seq++
	msg := tickMsg{Type: "tick", TreeID: s.treeID, Seq: s.seq, Nodes: s.pending}
	s.pending = make(map[string]string)
	s.mu.Unlock()

	s.broadcast(msg)
}

// PublishBlackboard broadcasts a snapshot of runtime variables to all connected
// clients as a "blackboard" message. Call it from user code after a tick, e.g.
//
//	tree.Tick(target, board)
//	dbg.PublishBlackboard(board.Dump(tree.GetID()))
//
// It is independent of the tick coalescing: send it as often (or as rarely) as
// makes sense for the variables being watched. Safe to call with no clients
// connected (it just no-ops the marshal/broadcast).
func (s *WSServer) PublishBlackboard(data map[string]interface{}) {
	s.mu.Lock()
	treeID := s.treeID
	s.mu.Unlock()
	s.broadcast(blackboardMsg{Type: "blackboard", TreeID: treeID, Data: data})
}

// broadcast marshals msg once and queues it to every client, dropping the frame
// for any client whose buffer is full.
func (s *WSServer) broadcast(msg interface{}) {
	data, err := jsonMarshal(msg)
	if err != nil {
		return
	}
	s.mu.Lock()
	for c := range s.clients {
		select {
		case c.send <- data:
		default: // client is slow; drop this frame
		}
	}
	s.mu.Unlock()
}

// --- connection handling ---

func (s *WSServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := s.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &client{conn: conn, send: make(chan []byte, sendBuffer)}

	s.mu.Lock()
	s.clients[c] = struct{}{}
	hello := helloMsg{Type: "hello", TreeID: s.treeID, NodeCount: len(s.last)}
	// Snapshot current statuses so a late-joining client sees live state
	// immediately, not only nodes that change after it connects.
	snapshot := tickMsg{Type: "tick", TreeID: s.treeID, Seq: s.seq, Nodes: cloneMap(s.last)}
	s.mu.Unlock()

	if data, err := jsonMarshal(hello); err == nil {
		c.send <- data
	}
	if len(snapshot.Nodes) > 0 {
		if data, err := jsonMarshal(snapshot); err == nil {
			c.send <- data
		}
	}

	go s.writePump(c)
	s.readPump(c) // blocks until the client disconnects
}

// writePump drains c.send to the connection. It is the only writer for conn.
func (s *WSServer) writePump(c *client) {
	for data := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			return
		}
	}
}

// readPump discards inbound messages and detects disconnect. The control
// channel (editor → Go) is reserved for a future phase. On return the client is
// unregistered and its connection and send channel are closed.
func (s *WSServer) readPump(c *client) {
	defer s.removeClient(c)
	for {
		if _, _, err := c.conn.ReadMessage(); err != nil {
			return
		}
	}
}

func (s *WSServer) removeClient(c *client) {
	s.mu.Lock()
	if _, ok := s.clients[c]; ok {
		delete(s.clients, c)
		close(c.send)
	}
	s.mu.Unlock()
	_ = c.conn.Close()
}

// Close stops the flush loop and the HTTP server, and disconnects all clients.
// It is safe to call once.
func (s *WSServer) Close() error {
	if !s.closed.CompareAndSwap(false, true) {
		return nil
	}
	close(s.done)

	s.mu.Lock()
	for c := range s.clients {
		delete(s.clients, c)
		close(c.send)
		_ = c.conn.Close()
	}
	s.mu.Unlock()

	if s.httpServer != nil {
		return s.httpServer.Close()
	}
	return nil
}

// Addr returns the address the HTTP server is listening on, or "" if not
// started via NewWSServer.
func (s *WSServer) Addr() string {
	if s.httpServer == nil {
		return ""
	}
	return s.httpServer.Addr
}

// jsonMarshal is a thin wrapper so the marshal strategy lives in one place.
func jsonMarshal(v any) ([]byte, error) { return json.Marshal(v) }

// cloneMap returns a shallow copy so the caller can read a snapshot without
// holding the lock.
func cloneMap(m map[string]string) map[string]string {
	out := make(map[string]string, len(m))
	for k, v := range m {
		out[k] = v
	}
	return out
}
