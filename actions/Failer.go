package actions

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Failer struct {
	Action
}

func (this *Failer) OnTick(tick *Tick) bt.Status {
	return bt.FAILURE
}
