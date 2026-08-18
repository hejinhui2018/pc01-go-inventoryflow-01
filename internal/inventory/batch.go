package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sort"
)

type Batch struct {
	ID     string
	Events []domain.InventoryEvent
}

func (b Batch) Validate() error {
	if b.ID == "" {
		return errors.New("batch id is required")
	}
	if len(b.Events) == 0 {
		return errors.New("batch is empty")
	}
	for _, e := range b.Events {
		if err := e.Validate(); err != nil {
			return err
		}
	}
	return nil
}
func (b Batch) Ordered() []domain.InventoryEvent {
	out := append([]domain.InventoryEvent(nil), b.Events...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].Version < out[j].Version })
	return out
}
func (l *Ledger) AddBatch(b Batch) (int, error) {
	if err := b.Validate(); err != nil {
		return 0, err
	}
	n := 0
	for _, e := range b.Ordered() {
		if err := l.Add(e); err != nil {
			return n, err
		}
		n++
	}
	return n, nil
}
func Group(events []domain.InventoryEvent) map[string][]domain.InventoryEvent {
	out := map[string][]domain.InventoryEvent{}
	for _, e := range events {
		k := ledgerKey(e.SKU, e.Warehouse)
		out[k] = append(out[k], e)
	}
	return out
}
