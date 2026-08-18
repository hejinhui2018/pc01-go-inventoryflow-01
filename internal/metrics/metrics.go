package metrics

import "sync/atomic"

type Counters struct {
	events     atomic.Int64
	reconciles atomic.Int64
	errors     atomic.Int64
}

func (c *Counters) Event()     { c.events.Add(1) }
func (c *Counters) Reconcile() { c.reconciles.Add(1) }
func (c *Counters) Error()     { c.errors.Add(1) }
func (c *Counters) Snapshot() map[string]int64 {
	return map[string]int64{"events": c.events.Load(), "reconciles": c.reconciles.Load(), "errors": c.errors.Load()}
}
