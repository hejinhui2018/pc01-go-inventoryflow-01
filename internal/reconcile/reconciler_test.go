package reconcile

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"testing"
)

func TestApplyReceipt(t *testing.T) {
	r := New()
	out, err := r.Apply(domain.InventoryEvent{ID: "e1", SKU: "A", Warehouse: "W", Version: 1, Kind: domain.EventReceipt, Quantity: 10}, 0)
	if err != nil || out.State.OnHand != 10 {
		t.Fatalf("out=%+v err=%v", out, err)
	}
}
func TestReplaySameVersionIsIgnored(t *testing.T) {
	r := New()
	e := domain.InventoryEvent{ID: "e1", SKU: "A", Warehouse: "W", Version: 1, Kind: domain.EventReceipt, Quantity: 10}
	if _, err := r.Apply(e, 0); err != nil {
		t.Fatal(err)
	}
	out, err := r.Apply(e, 0)
	if err != nil {
		t.Fatal(err)
	}
	if out.State.OnHand != 10 || !out.Duplicate {
		t.Fatalf("replay changed state: %+v", out)
	}
}
func TestStaleVersionIsIgnored(t *testing.T) {
	r := New()
	_, _ = r.Apply(domain.InventoryEvent{ID: "e2", SKU: "A", Warehouse: "W", Version: 2, Kind: domain.EventReceipt, Quantity: 10}, 0)
	out, err := r.Apply(domain.InventoryEvent{ID: "e1", SKU: "A", Warehouse: "W", Version: 1, Kind: domain.EventReceipt, Quantity: 5}, 0)
	if err != nil || !out.Duplicate || out.State.OnHand != 10 {
		t.Fatalf("out=%+v err=%v", out, err)
	}
}
