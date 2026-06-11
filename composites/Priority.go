package composites

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Priority struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {bt.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Priority) OnTick(tick *Tick) bt.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != bt.FAILURE {
			return status
		}
	}
	return bt.FAILURE
}
