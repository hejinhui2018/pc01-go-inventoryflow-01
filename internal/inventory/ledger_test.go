package inventory

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"testing"
)

func TestLedgerOrdersEvents(t *testing.T) {
	l := NewLedger()
	_ = l.Add(domain.InventoryEvent{ID: "b", SKU: "A", Warehouse: "W", Version: 2, Kind: domain.EventReceipt, Quantity: 2})
	_ = l.Add(domain.InventoryEvent{ID: "a", SKU: "A", Warehouse: "W", Version: 1, Kind: domain.EventReceipt, Quantity: 1})
	if got := l.LatestVersion("A", "W"); got != 2 {
		t.Fatalf("latest=%d", got)
	}
	if got := l.Events("A", "W")[0].Version; got != 1 {
		t.Fatalf("first=%d", got)
	}
}
