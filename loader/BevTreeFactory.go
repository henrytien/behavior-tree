package loader

import (
	"fmt"

	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/actions"
	. "github.com/henrytien/behavior-tree/composites"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	. "github.com/henrytien/behavior-tree/decorators"
)

func createBaseStructMaps() *bt.RegisterStructMaps {
	st := bt.NewRegisterStructMaps()
	// Actions.
	st.Register("Error", &Error{})
	st.Register("Failer", &Failer{})
	st.Register("Runner", &Runner{})
	st.Register("Succeeder", &Succeeder{})
	st.Register("Wait", &Wait{})
	st.Register("RandWait", &RandWait{})
	st.Register("RandomSleep", &RandWait{}) // compatibility alias
	st.Register("Log", &Log{})
	// Composites.
	st.Register("MemPriority", &MemPriority{})
	st.Register("MemSequence", &MemSequence{})
	st.Register("Priority", &Priority{})
	st.Register("Sequence", &Sequence{})

	// Decorators.
	st.Register("Inverter", &Inverter{})
	st.Register("Limiter", &Limiter{})
	st.Register("Probability", &Probability{})
	st.Register("MaxTime", &MaxTime{})
	st.Register("Repeater", &Repeater{})
	st.Register("RepeatUntilFailure", &RepeatUntilFailure{})
	st.Register("RepeatUntilSuccess", &RepeatUntilSuccess{})
	return st
}

// CreateBehaviorTreeFromConfig creates a behavior tree from a configuration.
func CreateBehaviorTreeFromConfig(config *BTTreeCfg, extMap *bt.RegisterStructMaps) *BehaviorTree {
	baseMaps := createBaseStructMaps()
	tree := NewBehaviorTree()
	tree.Load(config, baseMaps, extMap)
	return tree
}

// CreateBevTreeFromConfig creates a behavior tree from a configuration.
//
// Deprecated: use CreateBehaviorTreeFromConfig.
func CreateBevTreeFromConfig(config *BTTreeCfg, extMap *bt.RegisterStructMaps) *BehaviorTree {
	return CreateBehaviorTreeFromConfig(config, extMap)
}

// CreateBehaviorTreeFromConfigSafe is the error-returning counterpart of
// CreateBehaviorTreeFromConfig. Loading an invalid configuration (an unknown
// node name, a missing or mistyped property, …) makes the underlying loader
// panic; this wrapper recovers from that panic and reports it as an error so
// callers can handle malformed input without crashing the process.
func CreateBehaviorTreeFromConfigSafe(config *BTTreeCfg, extMap *bt.RegisterStructMaps) (tree *BehaviorTree, err error) {
	defer func() {
		if r := recover(); r != nil {
			tree = nil
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	return CreateBehaviorTreeFromConfig(config, extMap), nil
}
