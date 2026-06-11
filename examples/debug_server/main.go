/*
Real-time debugging demo.

Loads the same tree as examples/load_from_tree, attaches the WebSocket debug
server, and ticks once per second forever. Connect the Behavior Tree Editor (or
any WebSocket client) to ws://localhost:6112/debug to watch node statuses
stream live. See behavior-tree-editor/docs/REALTIME_DEBUGGING.md.

	go run ./examples/debug_server
	# then, in another terminal:
	#   wscat -c ws://localhost:6112/debug
*/
package main

import (
	"fmt"
	"time"

	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	"github.com/henrytien/behavior-tree/debug"
	. "github.com/henrytien/behavior-tree/examples/share"
	. "github.com/henrytien/behavior-tree/loader"
)

func main() {
	treeConfig, ok := LoadTreeCfg("examples/load_from_tree/tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err — run from the repo root")
		return
	}

	maps := bt.NewRegisterStructMaps()
	maps.Register("Log", new(LogTest))

	tree := CreateBehaviorTreeFromConfig(treeConfig, maps)

	dbg := debug.NewWSServer(":6112")
	defer dbg.Close()
	tree.SetDebug(dbg)

	fmt.Println("debug server on ws://localhost:6112/debug — connect the editor, then watch nodes tick")

	board := NewBlackboard()
	for {
		tree.Tick(0, board)
		time.Sleep(time.Second)
	}
}
