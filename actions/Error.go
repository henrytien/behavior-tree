package actions

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Error struct {
	Action
}

func (this *Error) OnTick(tick *Tick) bt.Status {
	return bt.ERROR
}
