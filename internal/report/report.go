package report

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sort"
	"time"
)

type Row struct {
	SKU       string    `json:"sku"`
	Warehouse string    `json:"warehouse"`
	OnHand    int64     `json:"on_hand"`
	Reserved  int64     `json:"reserved"`
	Available int64     `json:"available"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Rows(states []domain.ReconcileState) []Row {
	out := make([]Row, 0, len(states))
	for _, s := range states {
		out = append(out, Row{SKU: s.SKU, Warehouse: s.Warehouse, OnHand: s.OnHand, Reserved: s.Reserved, Available: s.Available(), UpdatedAt: s.UpdatedAt})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].SKU+out[i].Warehouse < out[j].SKU+out[j].Warehouse })
	return out
}
