package actions

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Runner struct {
	Action
}

func (this *Runner) OnTick(tick *Tick) bt.Status {
	return bt.RUNNING
}
