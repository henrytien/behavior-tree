package loader

import (
	"testing"

	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
)

// An unknown node name makes the loader panic. The safe wrapper must turn that
// into an error instead of crashing the test process.
func TestCreateBehaviorTreeFromConfigSafe_InvalidNode(t *testing.T) {
	cfg := &config.BTTreeCfg{
		Title: "bad tree",
		Root:  "n1",
		Nodes: map[string]config.BTNodeCfg{
			"n1": {
				Id:       "n1",
				Name:     "ThisNodeTypeDoesNotExist",
				Category: "action",
				Title:    "broken",
			},
		},
	}

	tree, err := CreateBehaviorTreeFromConfigSafe(cfg, b3.NewRegisterStructMaps())
	if err == nil {
		t.Fatal("expected an error for an unknown node name, got nil")
	}
	if tree != nil {
		t.Fatalf("expected nil tree on error, got %v", tree)
	}
}

// A valid configuration must load without error through the safe wrapper.
func TestCreateBehaviorTreeFromConfigSafe_Valid(t *testing.T) {
	cfg := &config.BTTreeCfg{
		Title: "ok tree",
		Root:  "n1",
		Nodes: map[string]config.BTNodeCfg{
			"n1": {
				Id:       "n1",
				Name:     "Succeeder",
				Category: "action",
				Title:    "ok",
			},
		},
	}

	tree, err := CreateBehaviorTreeFromConfigSafe(cfg, b3.NewRegisterStructMaps())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tree == nil {
		t.Fatal("expected a tree, got nil")
	}
}
