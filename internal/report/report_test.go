package report

import (
	"bytes"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"testing"
)

func TestRowsAndSummary(t *testing.T) {
	rows := Rows([]domain.ReconcileState{{SKU: "B", Warehouse: "W", OnHand: 3, Reserved: 1}, {SKU: "A", Warehouse: "W", OnHand: 2, Reserved: 0}})
	if rows[0].SKU != "A" {
		t.Fatal(rows)
	}
	s := Summarize([]domain.ReconcileState{{SKU: "A", OnHand: 2, Reserved: 1}})
	if s.TotalOnHand != 2 || s.TotalReserved != 1 {
		t.Fatal(s)
	}
}
func TestWriteCSV(t *testing.T) {
	var b bytes.Buffer
	if err := WriteCSV(&b, []Row{{SKU: "A", Warehouse: "W", OnHand: 2}}); err != nil {
		t.Fatal(err)
	}
	if b.Len() == 0 {
		t.Fatal("empty csv")
	}
}
