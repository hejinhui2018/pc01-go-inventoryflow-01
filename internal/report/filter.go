package report

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"strings"
)

type Filter struct {
	SKU              string
	Warehouse        string
	MinimumAvailable *int64
	AlertsOnly       bool
}

func Match(s domain.ReconcileState, f Filter) bool {
	if f.SKU != "" && s.SKU != f.SKU {
		return false
	}
	if f.Warehouse != "" && s.Warehouse != f.Warehouse {
		return false
	}
	if f.MinimumAvailable != nil && s.Available() < *f.MinimumAvailable {
		return false
	}
	if f.AlertsOnly && s.Healthy() {
		return false
	}
	return true
}
func FilterStates(states []domain.ReconcileState, f Filter) []domain.ReconcileState {
	var out []domain.ReconcileState
	for _, s := range states {
		if Match(s, f) {
			out = append(out, s)
		}
	}
	return out
}
func SearchRows(rows []Row, query string) []Row {
	q := strings.ToLower(query)
	if q == "" {
		return rows
	}
	var out []Row
	for _, r := range rows {
		if strings.Contains(strings.ToLower(r.SKU), q) || strings.Contains(strings.ToLower(r.Warehouse), q) {
			out = append(out, r)
		}
	}
	return out
}
