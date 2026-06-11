package core

import (
	"fmt"

	bt "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
)

// BehaviorTree represents an executable behavior tree.
//
// A tree is usually loaded from a BTTreeCfg exported by the editor. Tick must
// be called periodically to propagate execution from the root node.
//
// BehaviorTree does not store per-target execution state. Runtime data such as
// open nodes and node call counts is stored in a Blackboard, which allows one
// tree instance to drive many targets with isolated state.
type BehaviorTree struct {
	id          string
	title       string
	description string
	properties  map[string]interface{}
	root        IBaseNode
	debug       interface{}
	dumpInfo    *config.BTTreeCfg
}

// NewBehaviorTree creates an initialized behavior tree.
func NewBehaviorTree() *BehaviorTree {
	tree := &BehaviorTree{}
	tree.Initialize()
	return tree
}

// NewBeTree creates an initialized behavior tree.
//
// Deprecated: use NewBehaviorTree.
func NewBeTree() *BehaviorTree {
	return NewBehaviorTree()
}

// Initialize resets the tree to its default state.
func (tree *BehaviorTree) Initialize() {
	tree.id = bt.CreateUUID()
	tree.title = "The behavior tree"
	tree.description = "Default description"
	tree.properties = make(map[string]interface{})
	tree.root = nil
	tree.debug = nil
}

// GetID returns the tree ID.
func (tree *BehaviorTree) GetID() string {
	return tree.id
}

// GetTitle returns the tree title.
func (tree *BehaviorTree) GetTitle() string {
	return tree.title
}

// GetTitile returns the tree title.
//
// Deprecated: use GetTitle.
func (tree *BehaviorTree) GetTitile() string {
	return tree.GetTitle()
}

// SetDebug assigns a debug object used during ticks.
func (tree *BehaviorTree) SetDebug(debug interface{}) {
	tree.debug = debug
}

// GetRoot returns the tree root node.
func (tree *BehaviorTree) GetRoot() IBaseNode {
	return tree.root
}

// Load populates the tree from a behavior tree configuration.
//
// data must follow the editor-compatible Behavior3 JSON model. Built-in nodes
// are resolved from maps, while custom nodes can be supplied through extMaps
// using the node names present in data.
func (tree *BehaviorTree) Load(data *config.BTTreeCfg, maps *bt.RegisterStructMaps, extMaps *bt.RegisterStructMaps) {
	tree.title = data.Title
	tree.description = data.Description
	tree.properties = data.Properties
	tree.dumpInfo = data

	nodes := make(map[string]IBaseNode)

	// First create each node without wiring parent-child relationships.
	for id, s := range data.Nodes {
		spec := &s
		var node IBaseNode

		if spec.Category == "tree" {
			node = new(SubTree)
		} else if extMaps != nil && extMaps.CheckElem(spec.Name) {
			if customNode, err := extMaps.New(spec.Name); err == nil {
				node = customNode.(IBaseNode)
			}
		} else if baseNode, err := maps.New(spec.Name); err == nil {
			node = baseNode.(IBaseNode)
		}

		if node == nil {
			panic("BehaviorTree.Load: invalid node name:" + spec.Name + ", title:" + spec.Title)
		}

		node.Ctor()
		node.Initialize(spec)
		node.SetBaseNodeWorker(node.(IBaseWorker))
		nodes[id] = node
	}

	// Then connect children and decorator child references.
	for id, spec := range data.Nodes {
		node := nodes[id]

		if node.GetCategory() == bt.COMPOSITE && spec.Children != nil {
			for i := 0; i < len(spec.Children); i++ {
				childID := spec.Children[i]
				comp := node.(IComposite)
				comp.AddChild(nodes[childID])
			}
		} else if node.GetCategory() == bt.DECORATOR && len(spec.Child) > 0 {
			dec := node.(IDecorator)
			dec.SetChild(nodes[spec.Child])
		}
	}

	tree.root = nodes[data.Root]
}

// dump returns the configuration used to load this tree.
func (tree *BehaviorTree) dump() *config.BTTreeCfg {
	return tree.dumpInfo
}

// Tick propagates a tick through the tree, starting from the root.
//
// target is passed through to custom nodes. blackboard stores runtime execution
// state, including open nodes from previous ticks. Tick panics if blackboard is
// nil.
func (tree *BehaviorTree) Tick(target interface{}, blackboard *Blackboard) bt.Status {
	if blackboard == nil {
		panic("The blackboard parameter is obligatory and must be an instance of bt.Blackboard")
	}

	tick := NewTick()
	tick.debug = tree.debug
	tick.target = target
	tick.Blackboard = blackboard
	tick.tree = tree

	dbg, _ := tree.debug.(Debugger)
	if dbg != nil {
		dbg.OnTickStart(tree.id)
		// Fire on every return path, including the early return below.
		defer dbg.OnTickEnd(tree.id)
	}

	state := tree.root._execute(tick)

	lastOpenNodes := blackboard._getTreeData(tree.id).OpenNodes
	currOpenNodes := append([]IBaseNode(nil), tick._openNodes...)

	l := len(lastOpenNodes)
	if l == len(currOpenNodes) {
		if l == 0 || lastOpenNodes[l-1] == currOpenNodes[l-1] {
			return state
		}
	}

	// Compute the close range for nodes left open by the previous tick.
	start := 0
	for i := 0; i < bt.MinInt(len(lastOpenNodes), len(currOpenNodes)); i++ {
		start = i
		if lastOpenNodes[i] != currOpenNodes[i] {
			break
		}
	}

	for i := len(lastOpenNodes) - 1; i >= start; i-- {
		lastOpenNodes[i]._close(tick)
	}

	blackboard._getTreeData(tree.id).OpenNodes = currOpenNodes
	blackboard.SetTree("nodeCount", tick._nodeCount, tree.id)

	return state
}

// Print writes the tree structure to stdout.
func (tree *BehaviorTree) Print() {
	printNode(tree.root, 0)
}

func printNode(root IBaseNode, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}

	fmt.Print("|—", root.GetTitle())

	if root.GetCategory() == bt.DECORATOR {
		dec := root.(IDecorator)
		if dec.GetChild() != nil {
			printNode(dec.GetChild(), indent+3)
		}
	}

	fmt.Println("")
	if root.GetCategory() == bt.COMPOSITE {
		comp := root.(IComposite)
		for i := 0; i < comp.GetChildCount(); i++ {
			printNode(comp.GetChild(i), indent+3)
		}
	}
}
