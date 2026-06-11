package actions

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) OnTick(tick *Tick) bt.Status {
	return bt.SUCCESS
}
