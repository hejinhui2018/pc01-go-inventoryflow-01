package workflow

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/store"
	"testing"
)

func TestReceiveAndReport(t *testing.T) {
	st, _ := store.NewFileStore(t.TempDir())
	s := NewService(st)
	if err := s.Receive(domain.InventoryEvent{ID: "e", SKU: "A", Warehouse: "W", Version: 1, Kind: domain.EventReceipt, Quantity: 8}); err != nil {
		t.Fatal(err)
	}
	_, sum := s.Report()
	if sum.TotalOnHand != 8 {
		t.Fatal(sum)
	}
}
