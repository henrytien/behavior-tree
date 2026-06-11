package loader

import (
	"testing"

	bt "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/core"
)

// nodeEvent records one OnNodeStatus call.
type nodeEvent struct {
	nodeID string
	status bt.Status
}

// recordingDebugger captures the full tick lifecycle for assertions.
type recordingDebugger struct {
	tickStarts []string
	tickEnds   []string
	nodes      []nodeEvent
	// order records the relative ordering of every callback, so a test can
	// assert OnTickStart precedes node events which precede OnTickEnd.
	order []string
}

func (d *recordingDebugger) OnTickStart(treeID string) {
	d.tickStarts = append(d.tickStarts, treeID)
	d.order = append(d.order, "start")
}

func (d *recordingDebugger) OnNodeStatus(treeID, nodeID string, status bt.Status) {
	d.nodes = append(d.nodes, nodeEvent{nodeID, status})
	d.order = append(d.order, "node")
}

func (d *recordingDebugger) OnTickEnd(treeID string) {
	d.tickEnds = append(d.tickEnds, treeID)
	d.order = append(d.order, "end")
}

// statusOf returns the recorded status for nodeID, or 0 if it was never seen.
func (d *recordingDebugger) statusOf(nodeID string) bt.Status {
	for _, e := range d.nodes {
		if e.nodeID == nodeID {
			return e.status
		}
	}
	return 0
}

func TestDebugger_ReportsLifecycleAndStatuses(t *testing.T) {
	tree := buildTree(t, "seq",
		node("seq", "Sequence", "composite", "a", "b"),
		node("a", "SuccessLeaf", "action"),
		node("b", "FailLeaf", "action"),
	)

	dbg := &recordingDebugger{}
	tree.SetDebug(dbg)

	if got := tree.Tick(0, core.NewBlackboard()); got != bt.FAILURE {
		t.Fatalf("tree status = %v, want FAILURE", got)
	}

	// Lifecycle: exactly one start and one end, both carrying the tree id.
	if len(dbg.tickStarts) != 1 || len(dbg.tickEnds) != 1 {
		t.Fatalf("want 1 start + 1 end, got %d start / %d end", len(dbg.tickStarts), len(dbg.tickEnds))
	}
	if dbg.tickStarts[0] != tree.GetID() || dbg.tickEnds[0] != tree.GetID() {
		t.Fatalf("lifecycle carried wrong tree id: start=%q end=%q want=%q",
			dbg.tickStarts[0], dbg.tickEnds[0], tree.GetID())
	}

	// Per-node statuses, keyed by the exported node id the editor uses.
	if got := dbg.statusOf("a"); got != bt.SUCCESS {
		t.Fatalf("node a status = %v, want SUCCESS", got)
	}
	if got := dbg.statusOf("b"); got != bt.FAILURE {
		t.Fatalf("node b status = %v, want FAILURE", got)
	}
	// Sequence stops at the first failure, so it reports FAILURE overall.
	if got := dbg.statusOf("seq"); got != bt.FAILURE {
		t.Fatalf("node seq status = %v, want FAILURE", got)
	}

	// Ordering: start first, end last, at least one node event in between.
	if dbg.order[0] != "start" {
		t.Fatalf("first event = %q, want start", dbg.order[0])
	}
	if dbg.order[len(dbg.order)-1] != "end" {
		t.Fatalf("last event = %q, want end", dbg.order[len(dbg.order)-1])
	}
}

func TestDebugger_NotCalledWhenUnset(t *testing.T) {
	// A tree with no debugger must run exactly as before — this guards the
	// "zero cost when off" promise. We can't observe the nil path directly, so
	// we assert the tick still produces the correct status with debug unset.
	tree := buildTree(t, "seq",
		node("seq", "Sequence", "composite", "a"),
		node("a", "SuccessLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != bt.SUCCESS {
		t.Fatalf("tree without debugger = %v, want SUCCESS", got)
	}
}
