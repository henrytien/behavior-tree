---
id: getting-started
title: Getting Started
sidebar_position: 2
---

# Getting Started

## Install

Requires Go 1.23 or later.

```bash
go get github.com/henrytien/behavior-tree
```

Keep the module path and package name distinct: the module path is `github.com/henrytien/behavior-tree`, while the root package declaration should be `package behaviortree`. Go package names cannot contain hyphens, so there is no `package behavior-tree`. These examples use `b3` as an explicit import alias for readable references such as `b3.SUCCESS`.

## Load from a tree file

The most common usage is to design a tree in the [online editor](https://henrytien.github.io/behavior-tree-editor/), export it as JSON, then load and run it in Go.

```go
package main

import (
	"fmt"

	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	"github.com/henrytien/behavior-tree/loader"
)

// A custom Action node
type LogTest struct {
	Action
	info string
}

func (n *LogTest) Initialize(setting *config.BTNodeCfg) {
	n.Action.Initialize(setting)
	n.info = setting.GetPropertyAsString("info")
}

func (n *LogTest) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest:", n.info)
	return b3.SUCCESS
}

func main() {
	// 1. Load the exported tree config
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		panic("load tree failed")
	}

	// 2. Register custom nodes
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	// 3. Build the behavior tree
	tree := loader.CreateBehaviorTreeFromConfig(treeConfig, maps)
	tree.Print() // print the tree structure

	// 4. One blackboard per object; tick in a loop
	board := NewBlackboard()
	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}
}
```

## Key points

- **`Tick(target, blackboard)`**: `target` can be any object (usually the AI's owner). Built-in nodes never use it — it is only accessed by custom nodes.
- **`blackboard` is mandatory**: the tree stores no state; everything lives in the blackboard. One tree can drive many objects, each with its own blackboard.
- **Custom nodes** are registered via `RegisterStructMaps.Register(name, instance)`, where `name` matches the node's `name` field in the editor.

See [Examples](./examples) for more loading styles (from a project file, raw project file, subtrees, …).
