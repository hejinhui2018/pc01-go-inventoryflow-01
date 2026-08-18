package report

import (
	"fmt"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Alert struct {
	Key      string `json:"key"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

func Alerts(states []domain.ReconcileState) []Alert {
	var out []Alert
	for _, s := range states {
		if s.OnHand < 0 {
			out = append(out, Alert{Key: s.SKU + "@" + s.Warehouse, Severity: "critical", Message: "on-hand is negative"})
		}
		if s.Reserved > s.OnHand {
			out = append(out, Alert{Key: s.SKU + "@" + s.Warehouse, Severity: "warning", Message: fmt.Sprintf("reserved %d exceeds on-hand %d", s.Reserved, s.OnHand)})
		}
	}
	return out
}
