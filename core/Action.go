package core

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
)

type IAction interface {
	IBaseNode
}

/**
 * Action is the base class for all action nodes. Thus, if you want to create
 * new custom action nodes, you need to inherit from this class. For example,
 * take a look at the Runner action:
 *
 *     var Runner = bt.Class(bt.Action, {
 *       name: 'Runner',
 *
 *       tick: function(tick) {
 *         return bt.RUNNING;
 *       }
 *     });
 *
 * @module b3
 * @class Action
 * @extends BaseNode
**/
type Action struct {
	BaseNode
	BaseWorker
}

func (this *Action) Ctor() {
	this.category = bt.ACTION
}
func (this *Action) Initialize(params *BTNodeCfg) {

	//this.id = bt.CreateUUID()
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.parameters = make(map[string]interface{})
	this.properties = make(map[string]interface{})
}
