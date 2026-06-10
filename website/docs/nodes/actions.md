---
id: actions
title: 行为节点 Action
sidebar_position: 3
---

# 行为节点 Action

行为节点是叶子节点，执行具体逻辑并返回状态。下面是内建的行为节点；业务逻辑通常通过**自定义 Action** 实现（见[快速开始](../getting-started)）。

## Succeeder

总是返回 `SUCCESS`。

## Failer

总是返回 `FAILURE`。

## Error

总是返回 `ERROR`。

## Runner

总是返回 `RUNNING`。常用于占位或测试。

## Log

打印日志，返回 `SUCCESS`。

- 属性 `info`：要打印的字符串。

## Wait（等待）

等待固定时长后返回 `SUCCESS`，期间返回 `RUNNING`。

- 属性 `milliseconds`：等待的毫秒数。

在 `OnOpen` 记录起始时间，`OnTick` 中比较当前时间判断是否到期，因此跨帧保持同一截止时间。

## RandWait（随机等待）

等待一个**随机时长**后返回 `SUCCESS`，不阻塞 goroutine。随机时长在 `OnOpen` 时选定并存入黑板，`RUNNING` 期间保持同一截止时间。

支持的属性：

| 属性 | 说明 |
| --- | --- |
| `min_ms` / `max_ms` | 随机范围（毫秒），首选 |
| `timemini` / `timemax` | behavior3editor 的旧别名 |
| `milliseconds` | 固定时长回退，兼容 `Wait` |

> 注册时 `RandomSleep` 是 `RandWait` 的兼容别名。

| 注册名 | 行为 |
| --- | --- |
| `Succeeder` | 永远 SUCCESS |
| `Failer` | 永远 FAILURE |
| `Error` | 永远 ERROR |
| `Runner` | 永远 RUNNING |
| `Log` | 打印日志 |
| `Wait` | 固定等待 |
| `RandWait` / `RandomSleep` | 随机等待 |
