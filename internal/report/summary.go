package report

import "github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"

type Summary struct {
	Items         int   `json:"items"`
	TotalOnHand   int64 `json:"total_on_hand"`
	TotalReserved int64 `json:"total_reserved"`
	Alerts        int   `json:"alerts"`
}

func Summarize(states []domain.ReconcileState) Summary {
	var s Summary
	s.Items = len(states)
	for _, item := range states {
		s.TotalOnHand += item.OnHand
		s.TotalReserved += item.Reserved
		if !item.Healthy() {
			s.Alerts++
		}
	}
	return s
}
