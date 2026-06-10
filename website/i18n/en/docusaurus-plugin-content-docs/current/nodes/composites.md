---
id: composites
title: Composite Nodes
sidebar_position: 1
---

# Composite Nodes

A composite node has multiple children and controls their execution order and combination logic.

Registration names are in `loader/BevTreeFactory.go`.

## Sequence

Runs children one after another:

- Any child returns `FAILURE` → return `FAILURE` immediately.
- Any child returns `RUNNING` → return `RUNNING`.
- All succeed → return `SUCCESS`.

**No memory**: every tick restarts from the first child.

## Priority (Selector)

Tries children in order — an "OR" logic:

- Any child returns `SUCCESS` → return `SUCCESS` immediately.
- Any child returns `RUNNING` → return `RUNNING`.
- All fail → return `FAILURE`.

**No memory**: every tick re-evaluates from the first child.

## MemSequence

Same as `Sequence` but **with memory**: once a child returns `RUNNING`, the next frame resumes from that child instead of restarting. Good for step-by-step flows where you don't want to repeat earlier steps.

## MemPriority

Same as `Priority` but **with memory**: the child that returned `RUNNING` is remembered, and the next frame resumes evaluation from there.

## Sequence or MemSequence?

| Scenario | Recommended |
| --- | --- |
| Re-check preconditions every frame | `Sequence` / `Priority` |
| Steps advance once and should not be redone | `MemSequence` / `MemPriority` |

| Name | Type |
| --- | --- |
| `Sequence` | Sequence |
| `Priority` | Priority |
| `MemSequence` | Memory sequence |
| `MemPriority` | Memory priority |
