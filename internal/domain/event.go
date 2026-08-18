package domain

import (
	"errors"
	"time"
)

type EventKind string

const (
	EventReceipt    EventKind = "receipt"
	EventShipment   EventKind = "shipment"
	EventAdjustment EventKind = "adjustment"
)

type InventoryEvent struct {
	ID         string    `json:"id"`
	SKU        string    `json:"sku"`
	Warehouse  string    `json:"warehouse"`
	Version    int64     `json:"version"`
	Kind       EventKind `json:"kind"`
	Quantity   int64     `json:"quantity"`
	OccurredAt time.Time `json:"occurred_at"`
	Reference  string    `json:"reference"`
}

func (e InventoryEvent) Validate() error {
	if e.ID == "" || e.SKU == "" || e.Warehouse == "" {
		return errors.New("event identity is required")
	}
	if e.Version <= 0 {
		return errors.New("event version must be positive")
	}
	if e.Quantity == 0 {
		return errors.New("event quantity cannot be zero")
	}
	if e.Kind != EventReceipt && e.Kind != EventShipment && e.Kind != EventAdjustment {
		return errors.New("unsupported event kind")
	}
	return nil
}

func (e InventoryEvent) Delta() int64 {
	if e.Kind == EventShipment {
		return -abs(e.Quantity)
	}
	return e.Quantity
}
func abs(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}
