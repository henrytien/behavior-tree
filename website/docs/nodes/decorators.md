---
id: decorators
title: 装饰节点 Decorator
sidebar_position: 2
---

# 装饰节点 Decorator

装饰节点只有**一个**子节点，用于修饰或控制该子节点的行为。

## Inverter（反转）

反转子节点的结果：`SUCCESS` ↔ `FAILURE`，`RUNNING` 保持不变。

## Repeater（重复）

重复 Tick 子节点，直到达到设定的重复次数，或子节点返回 `RUNNING` / `ERROR`。

- 属性 `maxLoop`：最大重复次数；为负数时表示无限重复。

## RepeatUntilSuccess（重复至成功）

反复执行子节点，直到其返回 `SUCCESS`（或达到 `maxLoop` 上限）。

## RepeatUntilFailure（重复至失败）

反复执行子节点，直到其返回 `FAILURE`（或达到 `maxLoop` 上限）。

## MaxTime（限时）

限制子节点的最长执行时间。超过 `maxTime`（毫秒）仍未结束则返回 `FAILURE`。

## Limiter（限次）

限制子节点最多被执行的次数。

- 属性 `maxLoop`：允许执行的最大次数，超出后返回 `FAILURE`。

## Probability（概率）

按概率决定是否执行子节点。决策在 `OnOpen` 时掷骰一次并存入黑板，避免在子节点返回 `RUNNING` 期间被打断。

支持的属性（按优先级）：

| 属性 | 说明 |
| --- | --- |
| `probability` | 0.0–1.0，首选 |
| `rate` | `probability` 的别名 |
| `percent` | 0–100 的百分比写法 |
| `skip_status` | 不执行时返回的状态：`success`（默认）/ `failure` / `error` |

示例：`probability = 0.3` 表示 30% 概率执行子节点，否则返回 `skip_status`。

| 注册名 | 作用 |
| --- | --- |
| `Inverter` | 反转结果 |
| `Repeater` | 重复 N 次 |
| `RepeatUntilSuccess` | 重复至成功 |
| `RepeatUntilFailure` | 重复至失败 |
| `MaxTime` | 限制执行时长 |
| `Limiter` | 限制执行次数 |
| `Probability` | 按概率执行 |
