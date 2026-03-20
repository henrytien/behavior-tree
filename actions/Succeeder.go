package actions

import (
	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Succeeder struct {
	Action
}

func (this *Succeeder) OnTick(tick *Tick) b3.Status {
	return b3.SUCCESS
}
