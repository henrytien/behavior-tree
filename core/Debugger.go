package core

import bt "github.com/henrytien/behavior-tree"

// Debugger receives node execution events during a tick.
//
// Attach an implementation with BehaviorTree.SetDebug. The tree then reports
// the tick lifecycle and each visited node's returned status, letting external
// tools — such as the editor's real-time debugger — visualize execution. See
// docs/REALTIME_DEBUGGING for the protocol this feeds.
//
// Implementations are called on the tick path, so they must not block: a slow
// debugger slows every tick. The WebSocket server in package debug coalesces
// status changes off the tick goroutine to keep this cheap.
//
// When no debugger is set, BehaviorTree.debug is nil and these methods are
// never called, so existing trees pay no cost.
type Debugger interface {
	// OnTickStart is called once at the start of each tick, before the root
	// node executes.
	OnTickStart(treeID string)

	// OnNodeStatus is called once for every node visited during the tick, with
	// the status the node returned. status is one of bt.RUNNING, bt.SUCCESS,
	// bt.FAILURE or bt.ERROR.
	OnNodeStatus(treeID, nodeID string, status bt.Status)

	// OnTickEnd is called once after the tick completes, after any deferred
	// closes from the previous tick have run.
	OnTickEnd(treeID string)
}
