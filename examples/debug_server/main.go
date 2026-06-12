/*
Real-time debugging demo.

Loads examples/debug_server/demo_tree.json — a looping MemSequence of Wait
nodes — and attaches the WebSocket debug server. Because each Wait stays
RUNNING for a couple of seconds, the editor shows nodes turn blue (running)
then green (success) one after another, looping forever: a visible "flow".

Connect the Behavior Tree Editor (or any WebSocket client) to
ws://localhost:6112/debug. Import the SAME demo_tree.json into the editor so
node ids match. See behavior-tree-editor/docs/REALTIME_DEBUGGING.md.

	go run ./examples/debug_server
*/
package main

import (
	"fmt"
	"time"

	bt "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	"github.com/henrytien/behavior-tree/debug"
	. "github.com/henrytien/behavior-tree/loader"
)

func main() {
	treeConfig, ok := LoadTreeCfg("examples/debug_server/demo_tree.json")
	if !ok {
		fmt.Println("LoadTreeCfg err — run from the repo root")
		return
	}

	// All node types in the demo tree (Wait, Repeater, MemSequence, Log) are
	// built-ins, so no custom registrations are needed.
	tree := CreateBehaviorTreeFromConfig(treeConfig, bt.NewRegisterStructMaps())

	dbg := debug.NewWSServer(":6112")
	defer dbg.Close()
	tree.SetDebug(dbg)

	fmt.Println("debug server on ws://localhost:6112/debug")
	fmt.Println("import examples/debug_server/demo_tree.json into the editor, then Debug > Connect")

	// Tick several times a second so each Wait is sampled as RUNNING multiple
	// times — the highlight stays lit while the node waits, giving a smooth
	// flow rather than a single flash.
	board := NewBlackboard()
	for {
		tree.Tick(0, board)
		time.Sleep(200 * time.Millisecond)
	}
}
