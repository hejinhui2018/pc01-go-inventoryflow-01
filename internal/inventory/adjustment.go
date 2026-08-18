package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"time"
)

type Adjustment struct {
	SKU       string
	Warehouse string
	Delta     int64
	Reference string
	At        time.Time
}

func (a Adjustment) Event(id string, version int64) domain.InventoryEvent {
	kind := domain.EventAdjustment
	return domain.InventoryEvent{ID: id, SKU: a.SKU, Warehouse: a.Warehouse, Version: version, Kind: kind, Quantity: a.Delta, OccurredAt: a.At, Reference: a.Reference}
}
func ValidateAdjustment(a Adjustment) error {
	if a.SKU == "" || a.Warehouse == "" {
		return errors.New("adjustment identity is required")
	}
	if a.Delta == 0 {
		return errors.New("adjustment delta cannot be zero")
	}
	return nil
}
