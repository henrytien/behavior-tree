---
id: intro
title: Introduction
sidebar_position: 1
slug: /intro
---

# Introduction to behavior-tree

`behavior-tree` is a Go implementation of the [behavior3](https://github.com/behavior3) behavior tree. You can author logic nodes directly in the official online editor — the data format is kept identical to the original (it was translated from the JS version).

> `behavior-tree` is the external repository and product brand for the Go module `github.com/henrytien/behavior-tree`. The Go package declaration is `package behaviortree`, not the hyphenated `package behavior-tree`; examples can continue to use `b3` as an explicit import alias. This repository is derived from `magicsea/behavior3go` and is now maintained independently.

## Naming policy

- Use `behavior-tree` for the external brand, repository name, documentation site, and module path.
- Use `behaviortree` for the Go package declaration because Go identifiers cannot contain hyphens.
- Documentation examples may continue to use `b3 "github.com/henrytien/behavior-tree"` to keep examples compact and compatible with behavior3-style constants.

## How it differs from a typical behavior tree

This behavior tree is a little different from the usual implementation:

- **Stateless structure**: only a single copy of the tree structure is kept; all runtime state lives in the **Blackboard**.
- A typical behavior tree needs one full tree per object to hold state. Here, a **single** tree instance can drive hundreds of objects, with each object's state isolated in its own blackboard.

This design is especially efficient for game servers that manage large numbers of AI entities.

## Node categories

A behavior tree is built from four kinds of nodes:

| Category | Constant | Description |
| --- | --- | --- |
| Composite | `composite` | Controls the execution order of children (Sequence / Priority, …) |
| Decorator | `decorator` | Modifies the behavior of a single child (Inverter / Repeater, …) |
| Action | `action` | Leaf node that runs concrete logic (Wait / Log, …) |
| Condition | `condition` | Leaf node that returns a boolean-like result |

Every `Tick` propagates a signal downward from the root. Each node returns one of four states: `SUCCESS`, `FAILURE`, `RUNNING`, `ERROR`.

## Next steps

- [Getting Started](./getting-started) — install and your first example
- [Concepts](./concepts) — Tick, Blackboard, node lifecycle
- [Node Reference](./nodes/composites) — built-in nodes in detail
