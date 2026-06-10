---
id: subtree
title: 子树 SubTree
sidebar_position: 5
---

# 子树 SubTree

`SubTree` 节点允许把一棵独立的行为树作为另一棵树中的一个节点来引用，从而拆分、复用复杂逻辑。

加载时，`category` 为 `"tree"` 的节点会被实例化为 `SubTree`（见 `core/BehaviorTree.go` 的 `Load`）。

## 使用前提

子树需要**专用的编辑器分支版本** `behavior3editor`，并在导出 node 时带上 `category` 字段。

## 工作方式

- 父树中的 `SubTree` 节点指向某个子树的 ID。
- Tick 传播到 `SubTree` 时，会进入对应子树执行，再把结果返回给父树。
- 各子树的记忆节点状态同样保存在黑板中。

## 已知问题

> **Q：子树的相同记忆节点的黑板信息是否会重复？**
>
> A：是个问题。但由于下一次进入会在 `OnOpen` 重置节点黑板信息，并不会产生错误表现。如果遇到 BUG 表现欢迎反馈。

## 示例

仓库中提供了子树相关示例：

- `examples/subtree` — 子树的使用示例（需专用编辑器分支版本）。
- `examples/memsubtree` — 记忆子树示例。

详见[示例](../examples)。
