package actions

import (
	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

type Runner struct {
	Action
}

func (this *Runner) OnTick(tick *Tick) b3.Status {
	return b3.RUNNING
}
