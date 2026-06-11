# behavior-tree

Go behavior tree implementation derived from https://github.com/behavior3

[![Go](https://github.com/henrytien/behavior-tree/actions/workflows/go.yml/badge.svg)](https://github.com/henrytien/behavior-tree/actions/workflows/go.yml)
[![Docs](https://github.com/henrytien/behavior-tree/actions/workflows/docs.yml/badge.svg)](https://github.com/henrytien/behavior-tree/actions/workflows/docs.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/henrytien/behavior-tree.svg)](https://pkg.go.dev/github.com/henrytien/behavior-tree)

> `behavior-tree` is the external repository and product brand for the module `github.com/henrytien/behavior-tree`. The project is derived from `magicsea/behavior3go` and is now maintained independently.

📖 **文档站 / Documentation:** https://henrytien.github.io/behavior-tree/
✏️ **在线编辑器 / Online editor:** https://henrytien.github.io/behavior-tree-editor/

## Naming policy

- Repository, module path, and documentation brand: `behavior-tree` / `github.com/henrytien/behavior-tree`.
- Go package declaration: `package behaviortree`. Go package declarations cannot contain hyphens, so the package declaration is not `behavior-tree`.
- Examples may use the explicit import alias `b3` for readability and compatibility with behavior3-style constants.

## 简介

带在线编辑器的行为树，可使用[在线编辑器](https://henrytien.github.io/behavior-tree-editor/)编辑逻辑节点。
使用 js 版本翻译，保持和原版的编辑器数据格式一致。
此行为树和一般的行为树略有不同：行为树结构只保持一份**无状态**，状态记录在**黑板（Blackboard）**里。一般行为树每个对象需要一份完整的树结构来保存状态，而这里只需**一个**树实例就能驱动成百上千个对象，各对象状态彼此隔离在自己的黑板中——非常适合需要管理大量 AI 实体的游戏服务器。

## 安装 Installation

需要 Go 1.23 或更高版本 / Requires Go 1.23+.

```bash
go get github.com/henrytien/behavior-tree
```

## 快速开始 Quick Start

在[在线编辑器](https://henrytien.github.io/behavior-tree-editor/)中设计行为树并导出为 JSON，然后在 Go 中加载执行：

```go
package main

import (
	"fmt"

	b3 "github.com/henrytien/behavior-tree"
	"github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	"github.com/henrytien/behavior-tree/loader"
)

// 自定义 Action 节点 / a custom Action node
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
	// 1. 加载导出的树配置 / load the exported tree config
	treeConfig, ok := config.LoadTreeCfg("tree.json")
	if !ok {
		panic("load tree failed")
	}

	// 2. 注册自定义节点 / register custom nodes
	maps := b3.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	// 3. 构建行为树 / build the tree
	tree := loader.CreateBehaviorTreeFromConfig(treeConfig, maps)
	tree.Print()

	// 4. 每个对象一块黑板，循环 Tick / one blackboard per object, tick in a loop
	board := NewBlackboard()
	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}
}
```

> `CreateBevTreeFromConfig` 作为旧名仍保留，等价于 `CreateBehaviorTreeFromConfig`。

## 内建节点 Built-in Nodes

| 分类 Category | 节点 Nodes |
| --- | --- |
| 组合 Composite | `Sequence`、`Priority`、`MemSequence`、`MemPriority` |
| 装饰 Decorator | `Inverter`、`Repeater`、`RepeatUntilSuccess`、`RepeatUntilFailure`、`MaxTime`、`Limiter`、`Probability` |
| 行为 Action | `Succeeder`、`Failer`、`Error`、`Runner`、`Log`、`Wait`、`RandWait`（别名 `RandomSleep`） |

各节点的属性与用法详见[文档站 · 节点参考](https://henrytien.github.io/behavior-tree/docs/nodes/composites)。

## 示例 Examples

- [load_from_tree](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_tree)：从导出的树文件加载
- [load_from_project](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_project)：从导出的工程文件加载
- [load_from_rawproject](https://github.com/henrytien/behavior-tree/tree/master/examples/load_from_rawproject)：从原生工程文件加载
- [subtree](https://github.com/henrytien/behavior-tree/tree/master/examples/subtree)：子树的使用示例
- [memsubtree](https://github.com/henrytien/behavior-tree/tree/master/examples/memsubtree)：记忆子树示例
- [mmoarpg](examples/mmoarpg/zt.b3)：一个 mmoarpg 类型游戏的行为示例，可用桌面版编辑器打开查看（Projects → open project）

运行示例 / run an example：

```bash
cd examples/load_from_tree && go run main.go
```

## 完整示例 Full Example

io 类游戏示例：`h5game/server`。`bin/b3.json` 为行为树数据，在编辑器中新建任意工程，选择 **Project → Import → Tree as json** 导入树即可还原工程，如图。

![image](https://github.com/henrytien/behavior-tree/blob/master/b3_simple1.png)

## 更新 Changelog

* 升级到 Go 1.23，新增中英双语文档站与 CI/GitHub Pages 自动部署
* 添加子树支持 `SubTree` 节点，需要编辑器导出 node 的 category 字段
* 添加 `RandWait` 随机等待节点：支持 `min_ms/max_ms`，兼容 `timemini/timemax` 与固定 `milliseconds`
* 添加 `Probability` 概率装饰器：支持 `probability/rate/percent` 与 `skip_status`

## FAQ

完整 FAQ 见[文档站](https://henrytien.github.io/behavior-tree/docs/faq)。常见问题：

- **Q：Tick 里的 target 如何调用？**
  用在 AI 里，一般 `target` 就是这个 AI 的拥有者，拥有者持有 blackboard 成员。内建节点不使用它，仅供自定义节点访问。
- **Q：如何设计打断一个进行中（RUNNING）的状态？**
  参考 [issue #15](https://github.com/henrytien/behavior-tree/issues/15)。
- **Q：子树相同记忆节点的黑板信息是否会重复？**
  下一次进入会在 `OnOpen` 重置节点黑板信息，并不会有错误表现。如遇 BUG 欢迎反馈。

## TODO

- [ ] 参数类型化
- [ ] 参数支持传递黑板值，利用格式 `@变量名`
- [ ] 子树支持自定义参数传递

## 其他参考 References

- 节点概念与实践提示：[设计要点](https://henrytien.github.io/behavior-tree/docs/design-notes)
- 行为树基础知识：[Wikipedia: Behavior tree](https://en.wikipedia.org/wiki/Behavior_tree_(artificial_intelligence,_robotics_and_control))
- 其他人写的旧版 behavior3go 代码介绍：[CSDN 博客](https://blog.csdn.net/u013272009/article/details/77131226)
- behavior3 的 Erlang 实现（含行为树基础分享）：[behavior3erl](https://github.com/dong50252409/behavior3erl)

## 上线项目 Shipped Projects

* [丛林大作战](https://www.taptap.com/app/31608)
