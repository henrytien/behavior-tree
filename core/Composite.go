package core

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
)

type IComposite interface {
	IBaseNode
	GetChildCount() int
	GetChild(index int) IBaseNode
	AddChild(child IBaseNode)
}

type Composite struct {
	BaseNode
	BaseWorker

	children []IBaseNode
}

func (this *Composite) Ctor() {

	this.category = bt.COMPOSITE
}

/**
 * Initialization method.
 *
 * @method Initialize
 * @construCtor
**/
func (this *Composite) Initialize(params *BTNodeCfg) {
	this.BaseNode.Initialize(params)
	//this.BaseNode.IBaseWorker = this
	this.children = make([]IBaseNode, 0)
	//fmt.Println("Composite Initialize")
}

/**
 *
 * @method GetChildCount
 * @getChildCount
**/
func (this *Composite) GetChildCount() int {
	return len(this.children)
}

// GetChild
func (this *Composite) GetChild(index int) IBaseNode {
	return this.children[index]
}

// AddChild
func (this *Composite) AddChild(child IBaseNode) {
	this.children = append(this.children, child)
}
func (this *Composite) tick(tick *Tick) bt.Status {
	bt.Logf("tick Composite1")
	return bt.ERROR
}
