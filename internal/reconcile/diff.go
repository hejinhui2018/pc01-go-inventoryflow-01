package reconcile

import "github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"

type Difference struct {
	SKU       string `json:"sku"`
	Warehouse string `json:"warehouse"`
	Expected  int64  `json:"expected"`
	Actual    int64  `json:"actual"`
	Delta     int64  `json:"delta"`
}

func Compare(expected, actual []domain.ReconcileState) []Difference {
	m := map[string]domain.ReconcileState{}
	for _, s := range actual {
		m[s.SKU+"@"+s.Warehouse] = s
	}
	var out []Difference
	for _, e := range expected {
		a := m[e.SKU+"@"+e.Warehouse]
		if e.OnHand != a.OnHand || e.Reserved != a.Reserved {
			out = append(out, Difference{SKU: e.SKU, Warehouse: e.Warehouse, Expected: e.OnHand, Actual: a.OnHand, Delta: e.OnHand - a.OnHand})
		}
	}
	return out
}
