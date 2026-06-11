package loader

import (
	"testing"

	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
	"github.com/henrytien/behavior-tree/core"
)

// The node registry instantiates nodes via reflect.New, which yields a
// zero-value struct — any field set on the registered prototype is lost. So
// the test leaves carry no per-instance config; instead each concrete type
// has a fixed return status and bumps a package-level tick counter, letting
// tests assert both the resulting status and how many times each leaf ran.
var (
	tickSuccess int
	tickFail    int
	tickRun     int
)

func resetTicks() { tickSuccess, tickFail, tickRun = 0, 0, 0 }

type successLeaf struct{ core.Action }

func (n *successLeaf) OnTick(*core.Tick) b3.Status { tickSuccess++; return b3.SUCCESS }

type failLeaf struct{ core.Action }

func (n *failLeaf) OnTick(*core.Tick) b3.Status { tickFail++; return b3.FAILURE }

type runLeaf struct{ core.Action }

func (n *runLeaf) OnTick(*core.Tick) b3.Status { tickRun++; return b3.RUNNING }

// node is a small helper for building a BTNodeCfg.
func node(id, name, category string, children ...string) config.BTNodeCfg {
	cfg := config.BTNodeCfg{Id: id, Name: name, Category: category, Title: name}
	if category == "composite" {
		cfg.Children = children
	} else if category == "decorator" && len(children) > 0 {
		cfg.Child = children[0]
	}
	return cfg
}

// buildTree constructs a tree from nodes and registers the test leaves.
func buildTree(t *testing.T, root string, nodes ...config.BTNodeCfg) *core.BehaviorTree {
	t.Helper()
	resetTicks()
	cfg := &config.BTTreeCfg{Title: "test", Root: root, Nodes: map[string]config.BTNodeCfg{}}
	for _, n := range nodes {
		cfg.Nodes[n.Id] = n
	}

	maps := b3.NewRegisterStructMaps()
	maps.Register("SuccessLeaf", &successLeaf{})
	maps.Register("FailLeaf", &failLeaf{})
	maps.Register("RunLeaf", &runLeaf{})

	tree, err := CreateBehaviorTreeFromConfigSafe(cfg, maps)
	if err != nil {
		t.Fatalf("failed to build tree: %v", err)
	}
	return tree
}

func TestSequence_AllSucceed(t *testing.T) {
	tree := buildTree(t, "seq",
		node("seq", "Sequence", "composite", "a", "b"),
		node("a", "SuccessLeaf", "action"),
		node("b", "SuccessLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.SUCCESS {
		t.Fatalf("Sequence with all-success children = %v, want SUCCESS", got)
	}
	if tickSuccess != 2 {
		t.Fatalf("expected both children ticked (2), got %d", tickSuccess)
	}
}

func TestSequence_StopsOnFailure(t *testing.T) {
	// First child fails -> sequence returns FAILURE and never ticks the second.
	tree := buildTree(t, "seq",
		node("seq", "Sequence", "composite", "a", "b"),
		node("a", "FailLeaf", "action"),
		node("b", "SuccessLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.FAILURE {
		t.Fatalf("Sequence with failing first child = %v, want FAILURE", got)
	}
	if tickFail != 1 {
		t.Fatalf("expected failing child ticked once, got %d", tickFail)
	}
	if tickSuccess != 0 {
		t.Fatalf("expected second child NOT ticked, got %d", tickSuccess)
	}
}

func TestPriority_PicksFirstNonFailure(t *testing.T) {
	// First child fails, second succeeds -> Priority returns SUCCESS.
	tree := buildTree(t, "pri",
		node("pri", "Priority", "composite", "a", "b"),
		node("a", "FailLeaf", "action"),
		node("b", "SuccessLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.SUCCESS {
		t.Fatalf("Priority = %v, want SUCCESS", got)
	}
	if tickFail != 1 || tickSuccess != 1 {
		t.Fatalf("expected both children ticked once, got F=%d S=%d", tickFail, tickSuccess)
	}
}

func TestPriority_AllFail(t *testing.T) {
	tree := buildTree(t, "pri",
		node("pri", "Priority", "composite", "a", "b"),
		node("a", "FailLeaf", "action"),
		node("b", "FailLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.FAILURE {
		t.Fatalf("Priority with all-fail children = %v, want FAILURE", got)
	}
}

func TestInverter_FlipsResult(t *testing.T) {
	tree := buildTree(t, "inv",
		node("inv", "Inverter", "decorator", "a"),
		node("a", "FailLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.SUCCESS {
		t.Fatalf("Inverter(FAILURE) = %v, want SUCCESS", got)
	}
}

func TestTickPropagation_RunningPropagates(t *testing.T) {
	// A running leaf inside a sequence makes the whole tree report RUNNING.
	tree := buildTree(t, "seq",
		node("seq", "Sequence", "composite", "a", "b"),
		node("a", "RunLeaf", "action"),
		node("b", "SuccessLeaf", "action"),
	)
	if got := tree.Tick(0, core.NewBlackboard()); got != b3.RUNNING {
		t.Fatalf("Sequence with running first child = %v, want RUNNING", got)
	}
	if tickRun != 1 {
		t.Fatalf("expected running child ticked once, got %d", tickRun)
	}
	if tickSuccess != 0 {
		t.Fatalf("expected second child NOT ticked while first is RUNNING, got %d", tickSuccess)
	}
}
