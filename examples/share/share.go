package share

import (
	"fmt"
	bt "github.com/henrytien/behavior-tree"
	//. "github.com/henrytien/behavior-tree/actions"
	//. "github.com/henrytien/behavior-tree/composites"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	//. "github.com/henrytien/behavior-tree/decorators"
)

// 自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *Tick) bt.Status {
	fmt.Println("logtest:", tick.GetLastSubTree(), this.info)
	return bt.SUCCESS
}
