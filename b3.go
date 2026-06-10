// Package behaviortree provides shared constants and helpers for the
// behavior-tree module.
//
// The module path is github.com/henrytien/behavior-tree. The Go package name
// is behaviortree because package names are identifiers and cannot contain
// hyphens. Examples often import this package as b3 to keep behavior3-style
// constants readable.
package behaviortree

// VERSION is the current behavior-tree package version.
const (
	VERSION = "0.2.0"

	// COMPOSITE identifies composite nodes.
	COMPOSITE = "composite"

	// DECORATOR identifies decorator nodes.
	DECORATOR = "decorator"

	// ACTION identifies action nodes.
	ACTION = "action"

	// CONDITION identifies condition nodes.
	CONDITION = "condition"
)

// Status is the return state produced by behavior tree nodes.
type Status uint8

const (
	// SUCCESS indicates that a node completed successfully.
	SUCCESS Status = 1

	// FAILURE indicates that a node completed unsuccessfully.
	FAILURE Status = 2

	// RUNNING indicates that a node is still executing.
	RUNNING Status = 3

	// ERROR indicates that a node failed with an execution error.
	ERROR Status = 4
)
