package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sort"
	"sync"
)

type Ledger struct {
	mu     sync.RWMutex
	events map[string][]domain.InventoryEvent
}

func NewLedger() *Ledger                     { return &Ledger{events: make(map[string][]domain.InventoryEvent)} }
func ledgerKey(sku, warehouse string) string { return sku + "@" + warehouse }
func (l *Ledger) Add(e domain.InventoryEvent) error {
	if err := e.Validate(); err != nil {
		return err
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	key := ledgerKey(e.SKU, e.Warehouse)
	for _, old := range l.events[key] {
		if old.ID == e.ID {
			return errors.New("event already exists")
		}
	}
	l.events[key] = append(l.events[key], e)
	sort.Slice(l.events[key], func(i, j int) bool { return l.events[key][i].Version < l.events[key][j].Version })
	return nil
}
func (l *Ledger) Events(sku, warehouse string) []domain.InventoryEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()
	src := l.events[ledgerKey(sku, warehouse)]
	return append([]domain.InventoryEvent(nil), src...)
}
func (l *Ledger) All() []domain.InventoryEvent {
	l.mu.RLock()
	defer l.mu.RUnlock()
	var out []domain.InventoryEvent
	for _, es := range l.events {
		out = append(out, es...)
	}
	return out
}
func (l *Ledger) LatestVersion(sku, warehouse string) int64 {
	es := l.Events(sku, warehouse)
	var v int64
	for _, e := range es {
		if e.Version > v {
			v = e.Version
		}
	}
	return v
}
