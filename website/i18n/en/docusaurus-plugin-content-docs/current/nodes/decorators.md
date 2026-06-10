---
id: decorators
title: Decorator Nodes
sidebar_position: 2
---

# Decorator Nodes

A decorator node has exactly **one** child and modifies or controls that child's behavior.

## Inverter

Inverts the child's result: `SUCCESS` ↔ `FAILURE`. `RUNNING` is unchanged.

## Repeater

Repeatedly ticks the child until a configured repeat count is reached, or the child returns `RUNNING` / `ERROR`.

- Property `maxLoop`: maximum repeats; a negative value means infinite.

## RepeatUntilSuccess

Repeatedly runs the child until it returns `SUCCESS` (or `maxLoop` is reached).

## RepeatUntilFailure

Repeatedly runs the child until it returns `FAILURE` (or `maxLoop` is reached).

## MaxTime

Caps the child's running time. If it has not finished after `maxTime` (milliseconds), returns `FAILURE`.

## Limiter

Caps how many times the child may run.

- Property `maxLoop`: the maximum number of executions; beyond it, returns `FAILURE`.

## Probability

Runs the child according to a probability. The decision is rolled once in `OnOpen` and stored in the blackboard, so a child that returns `RUNNING` is not interrupted on later ticks.

Supported properties (by precedence):

| Property | Description |
| --- | --- |
| `probability` | 0.0–1.0, preferred |
| `rate` | alias of `probability` |
| `percent` | 0–100 percentage form |
| `skip_status` | status returned when not running: `success` (default) / `failure` / `error` |

Example: `probability = 0.3` means a 30% chance to run the child, otherwise it returns `skip_status`.

| Name | Effect |
| --- | --- |
| `Inverter` | Invert result |
| `Repeater` | Repeat N times |
| `RepeatUntilSuccess` | Repeat until success |
| `RepeatUntilFailure` | Repeat until failure |
| `MaxTime` | Limit running time |
| `Limiter` | Limit run count |
| `Probability` | Run by probability |
