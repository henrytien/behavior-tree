---
title: behavior-tree
---

# behavior-tree

## Golang 行为树 · Behavior Tree for Go

一个基于 [behavior3](https://github.com/behavior3) 的 Go 行为树实现，可直接使用官方在线编辑器编辑逻辑节点，与原版编辑器数据格式保持一致。

A Go implementation of the [behavior3](https://github.com/behavior3) behavior tree, fully compatible with the official online editor's data format.

### 快速链接 / Quick Links

- [📖 文档 / Documentation](/docs/intro) — 概念、节点参考与示例
- [🚀 快速开始 / Getting Started](/docs/getting-started)
- [💻 GitHub Repository](https://github.com/henrytien/behavior-tree)

### 特性 / Features

- **无状态树结构 / Stateless tree** — 行为树结构只保持一份，状态记录在黑板（Blackboard）里，一棵树可驱动成百上千个对象。
- **编辑器兼容 / Editor compatible** — 使用 JS 版本翻译，与官方 behavior3 在线编辑器数据格式一致。
- **子树支持 / Subtree support** — 支持 `SubTree` 节点组合复用。
- **丰富的内建节点 / Rich built-in nodes** — Composite、Decorator、Action、Condition 四大类，含随机等待、概率装饰器等。

> `behavior-tree` 是仓库和产品对外品牌，对应 Go module `github.com/henrytien/behavior-tree`。Go 包声明使用 `package behaviortree`，示例可继续使用 `b3` 作为显式导入别名。本仓库源自 `magicsea/behavior3go`，现独立维护。
