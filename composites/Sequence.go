package composites

import (
	_ "fmt"

	bt "github.com/henrytien/behavior-tree"
	_ "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
)

type Sequence struct {
	Composite
}

/**
 * Tick method.
 * @method tick
 * @param {bt.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Sequence) OnTick(tick *Tick) bt.Status {
	//fmt.Println("tick Sequence :", this.GetTitle())
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != bt.SUCCESS {
			return status
		}
	}
	return bt.SUCCESS
}
