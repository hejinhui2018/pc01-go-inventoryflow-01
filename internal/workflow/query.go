package workflow

import (
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/report"
)

func (s *Service) Events(sku, warehouse string) []domain.InventoryEvent {
	return s.Ledger.Events(sku, warehouse)
}
func (s *Service) State(sku, warehouse string) (domain.ReconcileState, bool) {
	return s.Reconciler.State(sku, warehouse)
}
func (s *Service) Report() ([]report.Row, report.Summary) { return s.SnapshotReport() }
