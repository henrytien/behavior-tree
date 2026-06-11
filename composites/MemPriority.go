package composites

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type MemPriority struct {
	Composite
}

/**
 * Open method.
 * @method open
 * @param {bt.Tick} tick A tick instance.
**/
func (this *MemPriority) OnOpen(tick *Tick) {
	tick.Blackboard.Set("runningChild", 0, tick.GetTree().GetID(), this.GetID())
}

/**
 * Tick method.
 * @method tick
 * @param {bt.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *MemPriority) OnTick(tick *Tick) bt.Status {
	var child = tick.Blackboard.GetInt("runningChild", tick.GetTree().GetID(), this.GetID())
	for i := child; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)

		if status != bt.FAILURE {
			if status == bt.RUNNING {
				tick.Blackboard.Set("runningChild", i, tick.GetTree().GetID(), this.GetID())
			}

			return status
		}
	}
	return bt.FAILURE
}
