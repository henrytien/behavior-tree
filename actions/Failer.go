package actions

import (
	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Failer struct {
	Action
}

func (this *Failer) OnTick(tick *Tick) b3.Status {
	return b3.FAILURE
}
