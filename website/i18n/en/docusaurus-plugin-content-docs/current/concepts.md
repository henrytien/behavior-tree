---
id: concepts
title: Concepts
sidebar_position: 3
---

# Concepts

## BehaviorTree

`BehaviorTree` represents the structure of a tree. It only holds the connections between nodes and **stores no runtime state**, so a single tree instance can be shared by many objects.

There are two ways to build one:

- **Manual construction**: set the root node and its children directly.
- **Load from config**: build from exported JSON via `Load` / `CreateBevTreeFromConfig` (recommended).

## Tick

A `Tick` is a single "heartbeat". Calling `tree.Tick(target, blackboard)` creates a `Tick` object carrying `target`, `blackboard`, and a tree reference, then propagates down from the root:

```
tree.Tick → root._execute(tick) → child._execute → ...
```

After each tick, the tree automatically closes nodes that were still open in the previous frame but were not visited this frame (handling `OnClose` on state transitions).

## Return states

Each node's `OnTick` returns one of four states:

| State | Meaning |
| --- | --- |
| `SUCCESS` | Completed successfully |
| `FAILURE` | Failed |
| `RUNNING` | Still running; continue next frame |
| `ERROR` | An error occurred |

## Blackboard

The blackboard is the sole carrier of state. It stores data in layers — `treeID → nodeID → key` — which means:

- When one tree drives different objects, their states never interfere (one blackboard per object).
- Nodes can remember data across frames. For example, `Wait` records its start time in `OnOpen` and reads it in later `OnTick` calls to decide whether it has timed out.

Common methods:

```go
board := core.NewBlackboard()
board.Set("key", value, treeID, nodeID)
v := board.GetInt64("key", treeID, nodeID)
b := board.GetBool("key", treeID, nodeID)
```

## Node lifecycle

When a node is `_execute`d, it goes through:

1. **`OnEnter`** — entering the node.
2. **`OnOpen`** — called only when the node was not open in the previous frame (first entry or restart). Good for initialization, e.g. randomizing a duration or rolling a probability.
3. **`OnTick`** — runs the core logic each frame and returns a state.
4. **`OnClose`** — called when the node closes (returns non-RUNNING, or is interrupted by its parent).
5. **`OnExit`** — leaving the node.

> Put "compute once" logic in `OnOpen` (e.g. `RandWait`'s random duration, `Probability`'s dice roll) to avoid recomputing it every frame while `RUNNING`, which would cause jittery behavior.

## SubTree

The `SubTree` node lets you reuse a whole tree as a node inside another tree, making it easy to split up complex logic. See the [SubTree docs](./nodes/subtree).
