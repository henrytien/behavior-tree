---
id: conditions
title: Condition Nodes
sidebar_position: 4
---

# Condition Nodes

Condition nodes are leaf nodes used for checks: return `SUCCESS` when the condition holds, `FAILURE` when it does not. They are usually combined with `Sequence` / `Priority` as a precondition gate for some piece of logic.

This library ships no concrete business condition nodes (such checks depend heavily on the specific game/system), so conditions are generally **implemented by you**.

## Custom condition node

Embed `core.Condition` and implement `OnTick`:

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
	// Reach the AI owner via tick.target and do the real check
	if enemyWithin(tick.target, c.distance) {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
```

Register it and use it in the editor:

```go
maps := b3.NewRegisterStructMaps()
maps.Register("IsEnemyNear", new(IsEnemyNear))
```

> Reading properties in `Initialize` and accessing runtime context via `tick.target` in `OnTick` is the common pattern for both custom Actions and Conditions.
