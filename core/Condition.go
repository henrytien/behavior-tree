package core

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
)

type ICondition interface {
	IBaseNode
}

type Condition struct {
	BaseNode
	BaseWorker
}

func (this *Condition) Ctor() {

	this.category = bt.CONDITION
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Condition) Initialize(params *BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
}
