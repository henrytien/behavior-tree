/*
从导出的树文件加载
*/
package main

import (
	"fmt"
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	. "github.com/henrytien/behavior-tree/examples/share"
	. "github.com/henrytien/behavior-tree/loader"
)

func main() {
	treeConfig, ok := LoadTreeCfg("tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err")
		return
	}
	//自定义节点注册
	maps := bt.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	//载入
	tree := CreateBehaviorTreeFromConfig(treeConfig, maps)
	tree.Print()

	//输入板
	board := NewBlackboard()
	//循环每一帧
	for i := 0; i < 5; i++ {
		tree.Tick(i, board)
	}

}
