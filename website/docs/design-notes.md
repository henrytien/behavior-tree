---
id: design-notes
title: 设计要点
sidebar_position: 4
---

# 设计要点

本页整理行为树的一些核心概念与实践提示，并标注本库实现中需要注意的地方。

## 瞬时节点与持续节点

按执行时长，节点可分为两类：

- **瞬时节点**：一帧内完成，立即返回 `SUCCESS` / `FAILURE`（如打印日志、设置变量、条件判断）。
- **持续节点**：跨多帧执行，进行中返回 `RUNNING`，直到完成才返回成功或失败（如 `Wait`、移动到某点）。

`RUNNING` 是行为树区别于普通状态机的关键：它让一个动作可以"持续若干帧"，而父节点能据此决定是继续等待还是切换分支。

## 记忆（Mem）节点

带 `Mem` 前缀的组合节点（`MemSequence` / `MemPriority`）会**记住上一帧返回 `RUNNING` 的子节点下标**，下一帧直接从该下标继续，而不是从头重新评估。

- 适合"按步骤推进、不应回退重做"的流程。
- 注意：记忆只在节点保持打开期间有效。若某帧未再访问到它，节点会被关闭并在 `OnOpen` 重置记忆。

## 选择 Sequence 还是 MemSequence？

| 需求 | 推荐 |
| --- | --- |
| 每帧都要重新检查前置条件（例如"敌人是否还在"）| `Sequence` / `Priority` |
| 步骤一次性推进、前置步骤不应重复执行 | `MemSequence` / `MemPriority` |

## MaxTime 不是阻塞等待

`MaxTime` 限制子节点的最长执行时间，但**它不会抢占（中断）子节点本身的执行**——子节点必须是非抢占式的。它的做法是：照常 Tick 子节点，一旦累计时间超过 `maxTime`，就把结果改判为 `FAILURE` 返回。因此 `MaxTime` 适合"给一个持续动作设上限"，而不是"睡眠 N 毫秒"。

## Limiter 的生命周期与黑板重置

`Limiter` 限制子节点在**整个生命周期内**最多被执行 `maxLoop` 次，超过后直接返回 `FAILURE` 而不再执行子节点。计数存在黑板里。

> ⚠️ **常见坑**：计数不会自动清零。例如把"战斗中最多放 3 次大招"用 `Limiter` 实现，那么一场战斗结束后若不重置该对象黑板，下一场战斗里这个分支会**再也不会触发**。请在合适的时机（如战斗结束、对象复用前）重置黑板。

`maxLoop` 是必填参数，缺失会 panic。

## Condition 与 Action 的职责

- **Condition（条件）**：只做判断的只读节点，相当于 `if (cond)` 里的测试函数，返回 `SUCCESS`（成立）或 `FAILURE`（不成立），不应产生副作用。
- **Action（行为）**：执行具体逻辑的叶子节点，可以有副作用（移动、攻击、改黑板等）。

把"判断"和"执行"分离到 Condition 与 Action，是保持行为树可读、可复用的基础。

## 子树的独立记忆链

`SubTree` 加载的子树拥有**独立的记忆链**：子树内部记忆节点的状态按子树自己的上下文存储，不会和父树相互污染。详见[子树文档](./nodes/subtree)。

## 延伸阅读

- 论文：*Behavior Trees in Robotics and AI: An Introduction*
- [behaviac 概念文档](https://www.behaviac.com/concepts/)
- [behavior3 的 Erlang 实现（含基础分享）](https://github.com/dong50252409/behavior3erl)
