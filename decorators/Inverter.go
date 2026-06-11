package decorators

import (
	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/core"
)

/**
 * The Inverter decorator inverts the result of the child, returning `SUCCESS`
 * for `FAILURE` and `FAILURE` for `SUCCESS`.
 *
 * @module b3
 * @class Inverter
 * @extends Decorator
**/
type Inverter struct {
	Decorator
}

/**
 * Tick method.
 * @method tick
 * @param {bt.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Inverter) OnTick(tick *Tick) bt.Status {
	if this.GetChild() == nil {
		return bt.ERROR
	}

	var status = this.GetChild().Execute(tick)
	if status == bt.SUCCESS {
		status = bt.FAILURE
	} else if status == bt.FAILURE {
		status = bt.SUCCESS
	}

	return status
}
