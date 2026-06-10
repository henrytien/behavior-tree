---
id: actions
title: Action Nodes
sidebar_position: 3
---

# Action Nodes

Action nodes are leaf nodes that run concrete logic and return a state. Below are the built-in actions; business logic is usually implemented via **custom actions** (see [Getting Started](../getting-started)).

## Succeeder

Always returns `SUCCESS`.

## Failer

Always returns `FAILURE`.

## Error

Always returns `ERROR`.

## Runner

Always returns `RUNNING`. Useful as a placeholder or for testing.

## Log

Prints a log line and returns `SUCCESS`.

- Property `info`: the string to print.

## Wait

Returns `SUCCESS` after a fixed duration; returns `RUNNING` in the meantime.

- Property `milliseconds`: the wait time in milliseconds.

It records the start time in `OnOpen` and compares the current time in `OnTick`, so the deadline stays the same across frames.

## RandWait

Returns `SUCCESS` after a **random duration**, without blocking the goroutine. The random duration is chosen in `OnOpen` and stored in the blackboard, so the deadline stays the same while `RUNNING`.

Supported properties:

| Property | Description |
| --- | --- |
| `min_ms` / `max_ms` | random range in milliseconds, preferred |
| `timemini` / `timemax` | legacy behavior3editor aliases |
| `milliseconds` | fixed-duration fallback, compatible with `Wait` |

> At registration, `RandomSleep` is a compatibility alias of `RandWait`.

| Name | Behavior |
| --- | --- |
| `Succeeder` | Always SUCCESS |
| `Failer` | Always FAILURE |
| `Error` | Always ERROR |
| `Runner` | Always RUNNING |
| `Log` | Print a log line |
| `Wait` | Fixed wait |
| `RandWait` / `RandomSleep` | Random wait |
