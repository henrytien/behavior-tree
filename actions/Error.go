package actions

import (
	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Error struct {
	Action
}

func (this *Error) OnTick(tick *Tick) b3.Status {
	return b3.ERROR
}
