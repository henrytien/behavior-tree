---
id: getting-started
title: 快速开始
sidebar_position: 2
---

# 快速开始

## 安装

需要 Go 1.23 或更高版本。

```bash
go get github.com/henrytien/behavior-tree
```

## 从树文件加载

最常见的用法是在[在线编辑器](https://henrytien.github.io/behavior-tree-editor/)中设计行为树，导出为 JSON，然后在 Go 中加载执行。

```go
package main

import (
	"fmt"

	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	"github.com/henrytien/behavior-tree/loader"
)

// 自定义 Action 节点
type LogTest struct {
	Action
	info string
}

func (n *LogTest) Initialize(setting *config.BTNodeCfg) {
	n.Action.Initialize(setting)
	n.info = setting.GetPropertyAsString("info")
}

func (n *LogTest) OnTick(tick *Tick) b3.Status {
	fmt.Println("logtest:", n.info)
	return b3.SUCCESS
}

func main() {
	// 1. 加载导出的树配置
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		panic("load tree failed")
	}

	// 2. 注册自定义节点
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	// 3. 构建行为树
	tree := loader.CreateBevTreeFromConfig(treeConfig, maps)
	tree.Print() // 打印树结构

	// 4. 每个对象一块黑板，循环 Tick
	board := NewBlackboard()
	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}
}
```

## 关键点

- **`Tick(target, blackboard)`**：`target` 可以是任意对象（一般是 AI 的拥有者），内建节点不使用它，仅供自定义节点访问。
- **`blackboard` 是必须的**：树本身不存状态，状态全在黑板里。一棵树可以配多块黑板驱动多个对象。
- **自定义节点**通过 `RegisterStructMaps.Register(name, instance)` 注册，`name` 对应编辑器里节点的 `name` 字段。

更多加载方式见[示例](./examples)（从工程文件、原生工程文件、子树等加载）。
