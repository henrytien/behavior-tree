---
id: concepts
title: 核心概念
sidebar_position: 3
---

# 核心概念

## BehaviorTree

`BehaviorTree` 表示一棵行为树的结构。它只持有节点的连接关系，**不存储任何运行状态**，因此一个树实例可以被多个对象共享。

构造方式有两种：

- **手动构造**：直接设置根节点和子节点。
- **从配置加载**：通过 `Load` / `CreateBehaviorTreeFromConfig` 从导出的 JSON 数据构建（推荐）。

## Tick

`Tick` 是一次"心跳"。调用 `tree.Tick(target, blackboard)` 会创建一个 `Tick` 对象，携带 `target`、`blackboard` 和树引用，从根节点开始向下传播：

```
tree.Tick → root._execute(tick) → 子节点 _execute → ...
```

每次 Tick 结束后，树会自动关闭上一帧仍打开、但本帧未被访问到的节点（即处理状态切换时的 `OnClose`）。

## 返回状态

每个节点的 `OnTick` 返回四种状态之一：

| 状态 | 含义 |
| --- | --- |
| `SUCCESS` | 执行成功 |
| `FAILURE` | 执行失败 |
| `RUNNING` | 仍在执行中，下一帧继续 |
| `ERROR` | 出错 |

## Blackboard（黑板）

黑板是状态的唯一载体。它按 `树ID → 节点ID → key` 分层存储数据，使得：

- 同一棵树驱动不同对象时，状态互不干扰（每个对象一块黑板）。
- 节点可以跨帧记忆数据，例如 `Wait` 节点在 `OnOpen` 时记下起始时间，在后续 `OnTick` 中读取以判断是否超时。

常用方法：

```go
board := core.NewBlackboard()
board.Set("key", value, treeID, nodeID)
v := board.GetInt64("key", treeID, nodeID)
b := board.GetBool("key", treeID, nodeID)
```

## 节点生命周期

一个节点在被 `_execute` 时，会经历：

1. **`OnEnter`** — 进入节点。
2. **`OnOpen`** — 仅当节点上一帧未打开时调用（首次进入或重新开始）。适合做初始化，如随机化时长、概率判定。
3. **`OnTick`** — 每帧执行核心逻辑，返回状态。
4. **`OnClose`** — 节点关闭时调用（返回非 RUNNING，或被父节点中断）。
5. **`OnExit`** — 离开节点。

> 把"只需算一次"的逻辑放在 `OnOpen`（例如 `RandWait` 的随机时长、`Probability` 的概率掷骰），可以避免节点在 `RUNNING` 期间每帧重新计算导致行为抖动。

## 子树 SubTree

`SubTree` 节点允许把一棵树作为另一棵树的节点复用，便于拆分复杂逻辑。详见[子树文档](./nodes/subtree)。
