package reconcile

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sort"
)

func SortEvents(events []domain.InventoryEvent) []domain.InventoryEvent {
	out := append([]domain.InventoryEvent(nil), events...)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Version == out[j].Version {
			return out[i].ID < out[j].ID
		}
		return out[i].Version < out[j].Version
	})
	return out
}
func Pending(events []domain.InventoryEvent, applied int64) []domain.InventoryEvent {
	var out []domain.InventoryEvent
	for _, e := range SortEvents(events) {
		if e.Version > applied {
			out = append(out, e)
		}
	}
	return out
}
