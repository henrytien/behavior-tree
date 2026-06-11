package share

import (
	bt "github.com/henrytien/behavior-tree"
	//. "github.com/henrytien/behavior-tree/actions"
	//. "github.com/henrytien/behavior-tree/composites"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	//. "github.com/henrytien/behavior-tree/decorators"
)

// 自定义action节点
type SetValue struct {
	Action
	value int
	key   string
}

func (this *SetValue) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *SetValue) OnTick(tick *Tick) bt.Status {
	tick.Blackboard.SetMem(this.key, this.value)
	return bt.SUCCESS
}

// 自定义action节点
type IsValue struct {
	Condition
	value int
	key   string
}

func (this *IsValue) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *IsValue) OnTick(tick *Tick) bt.Status {
	v := tick.Blackboard.GetInt(this.key, "", "")
	if v == this.value {
		return bt.SUCCESS
	}
	return bt.FAILURE
}
