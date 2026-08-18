package workflow

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/inventory"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/reconcile"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/report"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/store"
	"time"
)

type Service struct {
	Store        *store.FileStore
	Catalog      *inventory.Catalog
	Ledger       *inventory.Ledger
	Reservations *inventory.ReservationBook
	Reconciler   *reconcile.Reconciler
}

func NewService(st *store.FileStore) *Service {
	return &Service{Store: st, Catalog: inventory.NewCatalog(), Ledger: inventory.NewLedger(), Reservations: inventory.NewReservationBook(), Reconciler: reconcile.New()}
}
func (s *Service) Receive(e domain.InventoryEvent) error {
	if e.Kind == "" {
		e.Kind = domain.EventReceipt
	}
	if e.OccurredAt.IsZero() {
		e.OccurredAt = time.Now().UTC()
	}
	if err := s.Ledger.Add(e); err != nil {
		return err
	}
	_, err := s.Reconciler.Apply(e, s.Reservations.Reserved(e.SKU, e.Warehouse, nil))
	return err
}
func (s *Service) Reserve(r domain.Reservation) error {
	if err := s.Reservations.Create(r); err != nil {
		return err
	}
	st, ok := s.Reconciler.State(r.SKU, "WH-SH-01")
	if ok && st.Available() < r.Quantity {
		_ = s.Reservations.Change(r.ID, domain.ReservationReleased)
		return errors.New("insufficient available inventory")
	}
	return nil
}
func (s *Service) Reconcile(e domain.InventoryEvent) (domain.ReconcileResult, error) {
	if err := s.Ledger.Add(e); err != nil {
		return domain.ReconcileResult{}, err
	}
	return s.Reconciler.Apply(e, s.Reservations.Reserved(e.SKU, e.Warehouse, nil))
}
func (s *Service) SnapshotReport() ([]report.Row, report.Summary) {
	states := s.Reconciler.States()
	return report.Rows(states), report.Summarize(states)
}
