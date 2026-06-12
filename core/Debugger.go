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

// BreakpointController is an OPTIONAL extension of Debugger. A debugger that
// also implements it can pause execution at a node: the tree calls
// OnNodeEnter before a node ticks, and if the call blocks (because a breakpoint
// is set on that node), the whole tick — and the goroutine that called
// BehaviorTree.Tick — freezes there until the controller returns. This is how
// "stop the instant a skill fires" works: the freeze happens before the node's
// OnTick runs, so the blackboard reflects the moment of the hit.
//
// Because it blocks the tick goroutine, only enable breakpoints for debugging
// or stress testing, never in a latency-sensitive production tick loop.
//
// The type assertion for this interface is skipped entirely when no debugger
// is set, so non-debug trees pay nothing.
type BreakpointController interface {
	// OnNodeEnter is called just before a node's OnTick runs. Implementations
	// block here while the node is paused on a breakpoint, and return once the
	// client issues continue/step (or the controller is disabled). It returns
	// immediately when no breakpoint applies.
	OnNodeEnter(treeID, nodeID string)
}
