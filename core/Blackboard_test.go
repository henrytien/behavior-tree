package core

import "testing"

func TestBlackboard_Dump_MergesGlobalAndTree(t *testing.T) {
	bb := NewBlackboard()
	const tree = "tree-1"

	// Global var, a per-tree var, and a per-tree var that shadows a global key.
	bb.SetMem("g", 1)
	bb.SetMem("shared", "global")
	bb.SetTree("t", 2, tree)
	bb.SetTree("shared", "tree", tree)

	// Per-node memory must NOT appear in the dump.
	bb.Set("nodeOnly", 99, tree, "node-x")

	dump := bb.Dump(tree)

	if dump["g"] != 1 {
		t.Fatalf("global g = %v, want 1", dump["g"])
	}
	if dump["t"] != 2 {
		t.Fatalf("tree t = %v, want 2", dump["t"])
	}
	if dump["shared"] != "tree" {
		t.Fatalf("shared = %v, want tree (per-tree shadows global)", dump["shared"])
	}
	if _, present := dump["nodeOnly"]; present {
		t.Fatalf("per-node var leaked into dump: %v", dump["nodeOnly"])
	}
}

func TestBlackboard_Dump_IsACopy(t *testing.T) {
	bb := NewBlackboard()
	bb.SetMem("k", 1)

	dump := bb.Dump("")
	dump["k"] = 999 // mutating the snapshot must not affect the blackboard

	if got := bb.GetMem("k"); got != 1 {
		t.Fatalf("blackboard k = %v, want 1 (dump must be a copy)", got)
	}
}
