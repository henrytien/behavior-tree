---
id: design-notes
title: Design Notes
sidebar_position: 4
---

# Design Notes

This page collects core behavior-tree concepts and practical tips, with notes on what to watch for in this library's implementation.

## Instant vs. continuous nodes

By execution duration, nodes fall into two kinds:

- **Instant nodes**: finish within a single frame and immediately return `SUCCESS` / `FAILURE` (e.g. logging, setting a variable, a condition check).
- **Continuous nodes**: run across multiple frames, returning `RUNNING` while in progress until they finally succeed or fail (e.g. `Wait`, "move to a point").

`RUNNING` is what sets a behavior tree apart from a plain state machine: it lets an action "last several frames", and the parent decides whether to keep waiting or switch branches.

## Memory (Mem) nodes

Composite nodes with the `Mem` prefix (`MemSequence` / `MemPriority`) **remember the index of the child that returned `RUNNING`** last frame, and resume from that index next frame instead of re-evaluating from the start.

- Good for "advance step by step, don't redo earlier steps" flows.
- Note: memory only holds while the node stays open. If a frame doesn't reach it again, the node closes and resets its memory in `OnOpen`.

## Sequence or MemSequence?

| Need | Recommended |
| --- | --- |
| Re-check preconditions every frame (e.g. "is the enemy still there?") | `Sequence` / `Priority` |
| Steps advance once and earlier steps should not re-run | `MemSequence` / `MemPriority` |

## MaxTime is not a blocking wait

`MaxTime` caps how long a child may run, but **it does not preempt (interrupt) the child itself** — the child must be non-preemptive. Instead, it ticks the child as usual, and once the elapsed time exceeds `maxTime`, it overrides the result to `FAILURE`. So `MaxTime` is for "put an upper bound on a continuous action", not "sleep for N milliseconds".

## Limiter's lifecycle and blackboard reset

`Limiter` caps the child to at most `maxLoop` executions **over its entire lifecycle**; beyond that it returns `FAILURE` without running the child. The count lives in the blackboard.

> ⚠️ **Common pitfall**: the count does not reset on its own. For example, if you model "cast the ultimate at most 3 times in a battle" with `Limiter`, then once a battle ends, failing to reset that object's blackboard means the branch will **never trigger again** in the next battle. Reset the blackboard at the right time (battle end, before reusing the object).

`maxLoop` is mandatory; omitting it panics.

## Responsibilities of Condition vs. Action

- **Condition**: a read-only check node, like the test in `if (cond)`. Returns `SUCCESS` (holds) or `FAILURE` (does not). It should have no side effects.
- **Action**: a leaf node that runs concrete logic and may have side effects (move, attack, write the blackboard, …).

Splitting "checking" and "doing" into Conditions and Actions is the basis for keeping a behavior tree readable and reusable.

## SubTrees have independent memory chains

A tree loaded via `SubTree` has its **own memory chain**: the state of memory nodes inside the subtree is stored against the subtree's own context and does not pollute, or get polluted by, the parent tree. See the [SubTree docs](./nodes/subtree).

## Further reading

- Paper: *Behavior Trees in Robotics and AI: An Introduction*
- [Wikipedia: Behavior tree (AI, robotics and control)](https://en.wikipedia.org/wiki/Behavior_tree_(artificial_intelligence,_robotics_and_control))
- [An Erlang implementation of behavior3 (with fundamentals)](https://github.com/dong50252409/behavior3erl)
