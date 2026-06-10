---
id: intro
title: 简介
sidebar_position: 1
slug: /intro
---

# behavior3go 简介

`behavior3go` 是一个基于 [behavior3](https://github.com/behavior3) 的 Go 行为树实现，可直接使用官方在线编辑器编辑逻辑节点，并保持与原版编辑器一致的数据格式（使用 JS 版本翻译而来）。

> 本仓库源自 `magicsea/behavior3go`，现以 `github.com/henrytien/behavior-tree` 独立维护。

## 与一般行为树的区别

此行为树和一般的行为树略有不同：

- **结构无状态**：行为树结构只保持一份，所有运行状态记录在**黑板（Blackboard）**里。
- 一般的行为树每个对象都需要一份完整的树结构来保存状态；而这里只需**一个**树实例就能驱动成百上千个对象，各对象的状态彼此隔离在自己的黑板中。

这种设计在游戏服务器等需要管理大量 AI 实体的场景下尤为高效。

## 节点分类

行为树由四类节点构成：

| 分类 | 常量 | 说明 |
| --- | --- | --- |
| 组合 Composite | `composite` | 控制子节点的执行顺序（Sequence / Priority 等）|
| 装饰 Decorator | `decorator` | 修饰单个子节点的行为（Inverter / Repeater 等）|
| 行为 Action | `action` | 叶子节点，执行具体逻辑（Wait / Log 等）|
| 条件 Condition | `condition` | 叶子节点，返回判断结果 |

每次 `Tick` 都会从根节点开始向下传播信号，节点返回四种状态之一：`SUCCESS`、`FAILURE`、`RUNNING`、`ERROR`。

## 下一步

- [快速开始](./getting-started) — 安装与第一个示例
- [核心概念](./concepts) — Tick、Blackboard、节点生命周期
- [节点参考](./nodes/composites) — 内建节点详解
