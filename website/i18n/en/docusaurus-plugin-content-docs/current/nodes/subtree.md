---
id: subtree
title: SubTree
sidebar_position: 5
---

# SubTree

The `SubTree` node lets you reference a standalone behavior tree as a node inside another tree, so you can split up and reuse complex logic.

During loading, a node whose `category` is `"tree"` is instantiated as a `SubTree` (see `Load` in `core/BehaviorTree.go`).

## Prerequisite

SubTrees require the **dedicated editor branch** `behavior3editor`, which exports nodes with a `category` field.

## How it works

- A `SubTree` node in the parent tree points to a child tree's ID.
- When the tick reaches the `SubTree`, it enters the corresponding child tree, runs it, and returns the result to the parent.
- Memory-node state of each child tree is likewise stored in the blackboard.

## Known issue

> **Q: Is the blackboard data of identical memory nodes in a subtree duplicated?**
>
> A: It is a known issue, but because the next entry resets the node's blackboard data in `OnOpen`, it does not manifest as an error. If you hit a real bug, please report it.

## Examples

The repository includes subtree-related examples:

- `examples/subtree` — subtree usage (requires the dedicated editor branch).
- `examples/memsubtree` — memory subtree example.

See [Examples](../examples).
