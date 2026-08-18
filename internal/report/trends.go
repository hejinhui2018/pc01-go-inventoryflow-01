package report

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sort"
	"time"
)

type Point struct {
	At       time.Time
	OnHand   int64
	Reserved int64
}

func Daily(states []domain.ReconcileState) map[string]int64 {
	out := map[string]int64{}
	for _, s := range states {
		day := s.UpdatedAt.UTC().Format("2006-01-02")
		out[day] += s.OnHand
	}
	return out
}
func OrderedDays(values map[string]int64) []string {
	out := make([]string, 0, len(values))
	for k := range values {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
func Change(previous, current domain.ReconcileState) Point {
	return Point{At: current.UpdatedAt, OnHand: current.OnHand - previous.OnHand, Reserved: current.Reserved - previous.Reserved}
}
