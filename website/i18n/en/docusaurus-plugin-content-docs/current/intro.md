---
id: intro
title: Introduction
sidebar_position: 1
slug: /intro
---

# Introduction to behavior3go

`behavior3go` is a Go implementation of the [behavior3](https://github.com/behavior3) behavior tree. You can author logic nodes directly in the official online editor — the data format is kept identical to the original (it was translated from the JS version).

> This repository is derived from `magicsea/behavior3go` and is now maintained independently as `github.com/henrytien/behavior-tree`.

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
