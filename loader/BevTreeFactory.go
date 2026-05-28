package loader

import (
	_ "fmt"
	_ "reflect"

	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/actions"
	. "github.com/henrytien/behavior-tree/composites"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	. "github.com/henrytien/behavior-tree/decorators"
)

func createBaseStructMaps() *b3.RegisterStructMaps {
	st := b3.NewRegisterStructMaps()
	//actions
	st.Register("Error", &Error{})
	st.Register("Failer", &Failer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	st.Register("RandWait", &RandWait{})
	st.Register("RandomSleep", &RandWait{}) // compatibility alias
	st.Register("Log", &Log{})
	//composites
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("Sequence", &Sequence{})

	//decorators
	st.Register("Inverter", &Inverter{})
	st.Register("Limiter", &Limiter{})
	st.Register("Probability", &Probability{})
	st.Register("MaxTime", &MaxTime{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	return st
}

func CreateBevTreeFromConfig(config *BTTreeCfg, extMap *b3.RegisterStructMaps) *BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := NewBeTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}
