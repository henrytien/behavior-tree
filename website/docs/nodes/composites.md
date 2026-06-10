---
id: composites
title: 组合节点 Composite
sidebar_position: 1
---

# 组合节点 Composite

组合节点拥有多个子节点，负责控制它们的执行顺序与组合逻辑。

注册名见 `loader/BevTreeFactory.go`。

## Sequence（顺序）

依次执行子节点：

- 任一子节点返回 `FAILURE` → 立即返回 `FAILURE`。
- 任一子节点返回 `RUNNING` → 返回 `RUNNING`。
- 全部成功 → 返回 `SUCCESS`。

**不记忆**：每次 Tick 都从第一个子节点重新开始。

## Priority（优先 / Selector）

依次尝试子节点，相当于"或"逻辑：

- 任一子节点返回 `SUCCESS` → 立即返回 `SUCCESS`。
- 任一子节点返回 `RUNNING` → 返回 `RUNNING`。
- 全部失败 → 返回 `FAILURE`。

**不记忆**：每次 Tick 都从第一个子节点重新评估。

## MemSequence（记忆顺序）

行为同 `Sequence`，但**带记忆**：当某个子节点返回 `RUNNING` 后，下一帧会直接从该子节点继续，而不是从头开始。适合按步骤推进、不希望重复执行前置步骤的流程。

## MemPriority（记忆优先）

行为同 `Priority`，但**带记忆**：返回 `RUNNING` 的子节点会被记住，下一帧从该处继续评估。

## 选择 Sequence 还是 MemSequence？

| 场景 | 推荐 |
| --- | --- |
| 每帧都要重新检查前置条件 | `Sequence` / `Priority` |
| 步骤是一次性推进、不应回退重做 | `MemSequence` / `MemPriority` |

| 注册名 | 类型 |
| --- | --- |
| `Sequence` | 顺序 |
| `Priority` | 优先 |
| `MemSequence` | 记忆顺序 |
| `MemPriority` | 记忆优先 |
