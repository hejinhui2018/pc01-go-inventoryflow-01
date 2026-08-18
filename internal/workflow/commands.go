package workflow

import (
	"fmt"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"time"
)

func Receipt(id, sku, warehouse string, qty, version int64) domain.InventoryEvent {
	return domain.InventoryEvent{ID: id, SKU: sku, Warehouse: warehouse, Version: version, Kind: domain.EventReceipt, Quantity: qty, OccurredAt: time.Now().UTC(), Reference: fmt.Sprintf("receipt:%s", id)}
}
func Shipment(id, sku, warehouse string, qty, version int64) domain.InventoryEvent {
	return domain.InventoryEvent{ID: id, SKU: sku, Warehouse: warehouse, Version: version, Kind: domain.EventShipment, Quantity: qty, OccurredAt: time.Now().UTC(), Reference: fmt.Sprintf("shipment:%s", id)}
}
func NewReservation(id, order, sku string, qty int64) domain.Reservation {
	return domain.Reservation{ID: id, OrderID: order, SKU: sku, Quantity: qty, Status: domain.ReservationOpen, CreatedAt: time.Now().UTC()}
}
