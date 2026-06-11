package loader

import (
	"fmt"
	"reflect"
	"testing"

	bt "github.com/henrytien/behavior-tree"
	//. "github.com/henrytien/behavior-tree/actions"
	//. "github.com/henrytien/behavior-tree/composites"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
	//. "github.com/henrytien/behavior-tree/decorators"
)

type Test struct {
	value string
}

func (test *Test) Print() {
	fmt.Println(test.value)
}

func TestExample(t *testing.T) {
	maps := createBaseStructMaps()
	if data, err := maps.New("Runner"); err != nil {
		t.Error("Error:", err, data)
	} else {
		t.Log(reflect.TypeOf(data))
	}

}

// /////////////////////加载事例///////////////////////////
// 自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick *Tick) bt.Status {
	fmt.Println("logtest:", this.info)
	return bt.SUCCESS
}

func TestLoadTree(t *testing.T) {
	treeConfig, ok := LoadTreeCfg("tree.json")
	if ok {
		//自定义节点注册
		maps := bt.NewRegisterStructMaps()
		maps.Register("Log", new(LogTest))

		//载入
		tree := CreateBehaviorTreeFromConfig(treeConfig, maps)
		tree.Print()

		//输入板
		board := NewBlackboard()
		//循环每一帧
		for i := 0; i < 5; i++ {
			tree.Tick(i, board)
		}
	} else {
		t.Error("LoadTreeCfg err")
	}

}
