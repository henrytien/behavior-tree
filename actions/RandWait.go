package actions

import (
	"hash/fnv"
	"math/rand"
	"sync/atomic"
	"time"

	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
)

/**
 * RandWait waits for a random duration without blocking the caller goroutine.
 *
 * Supported properties:
 * - min_ms / max_ms: preferred range in milliseconds.
 * - timemini / timemax: behavior3editor legacy aliases.
 * - milliseconds: fixed wait fallback, compatible with Wait.
 *
 * The random duration is chosen in OnOpen and stored in blackboard, so a
 * RUNNING wait keeps the same deadline across ticks.
 */
var randWaitSeedCounter int64

type RandWait struct {
	Action
	minMs int64
	maxMs int64
}

func (this *RandWait) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	fixed := getInt64PropertyDefault(setting, "milliseconds", 0)
	this.minMs = getInt64PropertyDefault(setting, "min_ms", getInt64PropertyDefault(setting, "timemini", fixed))
	this.maxMs = getInt64PropertyDefault(setting, "max_ms", getInt64PropertyDefault(setting, "timemax", this.minMs))
	if this.minMs < 0 {
		this.minMs = 0
	}
	if this.maxMs < this.minMs {
		this.maxMs = this.minMs
	}
}

func (this *RandWait) OnOpen(tick *Tick) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	duration := this.randomDurationMs()
	treeID := tick.GetTree().GetID()
	nodeID := this.GetID()
	tick.Blackboard.Set("startTime", startTime, treeID, nodeID)
	tick.Blackboard.Set("duration", duration, treeID, nodeID)
}

func (this *RandWait) OnTick(tick *Tick) b3.Status {
	currTime := time.Now().UnixNano() / int64(time.Millisecond)
	treeID := tick.GetTree().GetID()
	nodeID := this.GetID()
	startTime := tick.Blackboard.GetInt64("startTime", treeID, nodeID)
	duration := tick.Blackboard.GetInt64("duration", treeID, nodeID)
	if currTime-startTime >= duration {
		return b3.SUCCESS
	}
	return b3.RUNNING
}

func (this *RandWait) randomDurationMs() int64 {
	if this.maxMs <= this.minMs {
		return this.minMs
	}
	span := this.maxMs - this.minMs + 1
	seed := time.Now().UnixNano() + hashString64(this.GetID()) + atomic.AddInt64(&randWaitSeedCounter, 1)
	return this.minMs + rand.New(rand.NewSource(seed)).Int63n(span)
}

func getInt64PropertyDefault(setting *BTNodeCfg, name string, fallback int64) int64 {
	if setting == nil || setting.Properties == nil {
		return fallback
	}
	value, ok := setting.Properties[name]
	if !ok {
		return fallback
	}
	switch v := value.(type) {
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	}
	return fallback
}

func hashString64(s string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return int64(h.Sum64() & 0x7fffffffffffffff)
}
