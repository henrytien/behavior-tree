package debug

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	bt "github.com/henrytien/behavior-tree"
)

// newTestServer starts a WSServer behind httptest with a fast flush interval so
// tests don't wait 100ms per frame. Returns the server, the ws:// URL, and a
// cleanup func.
func newTestServer(t *testing.T) (*WSServer, string, func()) {
	t.Helper()
	s := newServer(5 * time.Millisecond)
	go s.flushLoop()

	mux := http.NewServeMux()
	mux.HandleFunc("/debug", s.handleWS)
	hs := httptest.NewServer(mux)

	url := "ws" + strings.TrimPrefix(hs.URL, "http") + "/debug"
	cleanup := func() {
		_ = s.Close()
		hs.Close()
	}
	return s, url, cleanup
}

// readJSON reads one message and unmarshals it into a generic map, failing the
// test on timeout or error.
func readJSON(t *testing.T, c *websocket.Conn) map[string]any {
	t.Helper()
	_ = c.SetReadDeadline(time.Now().Add(time.Second))
	_, data, err := c.ReadMessage()
	if err != nil {
		t.Fatalf("read message: %v", err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal %q: %v", data, err)
	}
	return m
}

func TestWSServer_HelloOnConnect(t *testing.T) {
	s, url, cleanup := newTestServer(t)
	defer cleanup()

	// Establish a tree id before connecting.
	s.OnTickStart("tree-1")

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	msg := readJSON(t, c)
	if msg["type"] != "hello" {
		t.Fatalf("first message type = %v, want hello", msg["type"])
	}
	if msg["treeId"] != "tree-1" {
		t.Fatalf("hello treeId = %v, want tree-1", msg["treeId"])
	}
}

func TestWSServer_StreamsStatusChanges(t *testing.T) {
	s, url, cleanup := newTestServer(t)
	defer cleanup()
	s.OnTickStart("tree-1")

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()

	// Consume the hello frame.
	if msg := readJSON(t, c); msg["type"] != "hello" {
		t.Fatalf("expected hello, got %v", msg["type"])
	}

	// Report a status change; it should arrive as a coalesced tick frame.
	s.OnNodeStatus("tree-1", "node-a", bt.RUNNING)

	msg := readJSON(t, c)
	if msg["type"] != "tick" {
		t.Fatalf("message type = %v, want tick", msg["type"])
	}
	nodes, ok := msg["nodes"].(map[string]any)
	if !ok {
		t.Fatalf("nodes field missing or wrong type: %v", msg["nodes"])
	}
	if nodes["node-a"] != "running" {
		t.Fatalf("node-a = %v, want running", nodes["node-a"])
	}
}

func TestWSServer_OnlySendsChangedNodes(t *testing.T) {
	s, url, cleanup := newTestServer(t)
	defer cleanup()
	s.OnTickStart("tree-1")

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()
	readJSON(t, c) // hello

	// First change for node-a.
	s.OnNodeStatus("tree-1", "node-a", bt.RUNNING)
	if nodes := readJSON(t, c)["nodes"].(map[string]any); nodes["node-a"] != "running" {
		t.Fatalf("first frame node-a = %v, want running", nodes["node-a"])
	}

	// Reporting the SAME status must not produce a frame; a new status for a
	// different node must. We report node-a=running again (no change) and
	// node-b=success (change) — the next frame should contain only node-b.
	s.OnNodeStatus("tree-1", "node-a", bt.RUNNING)
	s.OnNodeStatus("tree-1", "node-b", bt.SUCCESS)

	nodes := readJSON(t, c)["nodes"].(map[string]any)
	if _, present := nodes["node-a"]; present {
		t.Fatalf("unchanged node-a should not be resent, got %v", nodes)
	}
	if nodes["node-b"] != "success" {
		t.Fatalf("node-b = %v, want success", nodes["node-b"])
	}
}

func TestWSServer_SeqIncreases(t *testing.T) {
	s, url, cleanup := newTestServer(t)
	defer cleanup()
	s.OnTickStart("tree-1")

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()
	readJSON(t, c) // hello

	s.OnNodeStatus("tree-1", "n", bt.RUNNING)
	first := readJSON(t, c)
	s.OnNodeStatus("tree-1", "n", bt.SUCCESS)
	second := readJSON(t, c)

	seq1 := first["seq"].(float64)
	seq2 := second["seq"].(float64)
	if seq2 <= seq1 {
		t.Fatalf("seq did not increase: %v then %v", seq1, seq2)
	}
}

func TestWSServer_PublishBlackboard(t *testing.T) {
	s, url, cleanup := newTestServer(t)
	defer cleanup()
	s.OnTickStart("tree-1")

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	defer c.Close()
	readJSON(t, c) // hello

	s.PublishBlackboard(map[string]interface{}{"ticks": 7, "phase": "B"})

	msg := readJSON(t, c)
	if msg["type"] != "blackboard" {
		t.Fatalf("message type = %v, want blackboard", msg["type"])
	}
	data, ok := msg["data"].(map[string]any)
	if !ok {
		t.Fatalf("data field missing or wrong type: %v", msg["data"])
	}
	if data["phase"] != "B" {
		t.Fatalf("phase = %v, want B", data["phase"])
	}
	if data["ticks"].(float64) != 7 {
		t.Fatalf("ticks = %v, want 7", data["ticks"])
	}
}

func TestStatusString(t *testing.T) {
	cases := map[bt.Status]string{
		bt.SUCCESS: "success",
		bt.FAILURE: "failure",
		bt.RUNNING: "running",
		bt.ERROR:   "error",
	}
	for status, want := range cases {
		if got := statusString(status); got != want {
			t.Fatalf("statusString(%v) = %q, want %q", status, got, want)
		}
	}
}
