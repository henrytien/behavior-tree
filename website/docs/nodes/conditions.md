---
id: conditions
title: 条件节点 Condition
sidebar_position: 4
---

# 条件节点 Condition

条件节点是叶子节点，用于做判断：返回 `SUCCESS` 表示条件成立，`FAILURE` 表示不成立。它通常配合 `Sequence` / `Priority` 使用，作为执行某段逻辑的前置门槛。

本库没有内建具体的业务条件节点（业务判断高度依赖具体游戏/系统），因此条件节点一般由你**自定义实现**。

## 自定义条件节点

继承 `core.Condition`，实现 `OnTick`：

```go
package main

import (
	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
)

type IsEnemyNear struct {
	Condition
	distance float64
}

func (c *IsEnemyNear) Initialize(setting *config.BTNodeCfg) {
	c.Condition.Initialize(setting)
	c.distance = setting.GetProperty("distance")
}

func (c *IsEnemyNear) OnTick(tick *Tick) b3.Status {
	// 通过 tick.target 拿到 AI 拥有者，做实际判断
	if enemyWithin(tick.target, c.distance) {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
```

注册后即可在编辑器中使用：

```go
maps := b3.NewRegisterStructMaps()
maps.Register("IsEnemyNear", new(IsEnemyNear))
```

> `Initialize` 中读取属性，`OnTick` 中通过 `tick.target` 访问运行时上下文，是自定义 Action 与 Condition 的通用模式。
