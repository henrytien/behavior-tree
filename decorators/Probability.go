package decorators

import (
	"hash/fnv"
	"math/rand"
	"strings"
	"sync/atomic"
	"time"

	b3 "github.com/henrytien/behavior-tree"
	. "github.com/henrytien/behavior-tree/config"
	. "github.com/henrytien/behavior-tree/core"
)

/**
 * Probability executes its child according to a probability.
 *
 * Supported properties:
 * - probability: 0.0-1.0, preferred.
 * - rate: alias of probability.
 * - percent: 0-100 alias.
 * - skip_status: success|failure|error, default success.
 *
 * The decision is made once in OnOpen and stored in blackboard. This avoids
 * interrupting a child that returns RUNNING on later ticks.
 */
var probabilitySeedCounter int64

type Probability struct {
	Decorator
	probability float64
	skipStatus  b3.Status
}

func (this *Probability) Initialize(setting *BTNodeCfg) {
	this.Decorator.Initialize(setting)
	this.probability = getFloatPropertyDefault(setting, "probability", getFloatPropertyDefault(setting, "rate", -1))
	if this.probability < 0 {
		percent := getFloatPropertyDefault(setting, "percent", 100)
		this.probability = percent / 100.0
	}
	if this.probability < 0 {
		this.probability = 0
	}
	if this.probability > 1 {
		this.probability = 1
	}

	switch strings.ToLower(strings.TrimSpace(getStringPropertyDefault(setting, "skip_status", "success"))) {
	case "failure", "fail":
		this.skipStatus = b3.FAILURE
	case "error":
		this.skipStatus = b3.ERROR
	default:
		this.skipStatus = b3.SUCCESS
	}
}

func (this *Probability) OnOpen(tick *Tick) {
	shouldRun := this.probability >= 1
	if !shouldRun && this.probability > 0 {
		seed := time.Now().UnixNano() + hashString64(this.GetID()) + atomic.AddInt64(&probabilitySeedCounter, 1)
		shouldRun = rand.New(rand.NewSource(seed)).Float64() <= this.probability
	}
	tick.Blackboard.Set("shouldRun", shouldRun, tick.GetTree().GetID(), this.GetID())
}

func (this *Probability) OnTick(tick *Tick) b3.Status {
	if this.GetChild() == nil {
		return b3.ERROR
	}
	shouldRun := tick.Blackboard.GetBool("shouldRun", tick.GetTree().GetID(), this.GetID())
	if !shouldRun {
		return this.skipStatus
	}
	return this.GetChild().Execute(tick)
}

func getFloatPropertyDefault(setting *BTNodeCfg, name string, fallback float64) float64 {
	if setting == nil || setting.Properties == nil {
		return fallback
	}
	value, ok := setting.Properties[name]
	if !ok {
		return fallback
	}
	switch v := value.(type) {
	case int:
		return float64(v)
	case int32:
		return float64(v)
	case int64:
		return float64(v)
	case uint:
		return float64(v)
	case uint32:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return v
	}
	return fallback
}

func getStringPropertyDefault(setting *BTNodeCfg, name, fallback string) string {
	if setting == nil || setting.Properties == nil {
		return fallback
	}
	value, ok := setting.Properties[name]
	if !ok {
		return fallback
	}
	if s, ok := value.(string); ok {
		return s
	}
	return fallback
}

func hashString64(s string) int64 {
	h := fnv.New64a()
	_, _ = h.Write([]byte(s))
	return int64(h.Sum64() & 0x7fffffffffffffff)
}
