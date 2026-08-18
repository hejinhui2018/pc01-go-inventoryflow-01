package reconcile

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sync"
	"time"
)

type Reconciler struct {
	mu     sync.Mutex
	states map[string]domain.ReconcileState
}

func New() *Reconciler                   { return &Reconciler{states: make(map[string]domain.ReconcileState)} }
func key(e domain.InventoryEvent) string { return e.SKU + "@" + e.Warehouse }
func (r *Reconciler) Apply(e domain.InventoryEvent, reserved int64) (domain.ReconcileResult, error) {
	if err := e.Validate(); err != nil {
		return domain.ReconcileResult{}, err
	}
	if reserved < 0 {
		return domain.ReconcileResult{}, errors.New("reserved quantity cannot be negative")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	s := r.states[key(e)]
	if s.SKU == "" {
		s.SKU = e.SKU
		s.Warehouse = e.Warehouse
	}
	if e.Version <= s.AppliedVersion {
		return domain.ReconcileResult{State: s, Duplicate: true, Reason: "stale event"}, nil
	}
	s.OnHand += e.Delta()
	s.AppliedVersion = e.Version
	s.Reserved = reserved
	s.UpdatedAt = time.Now().UTC()
	if !s.Healthy() {
		return domain.ReconcileResult{State: s, Applied: true, Reason: "inventory invariant violated"}, errors.New("inventory invariant violated")
	}
	r.states[key(e)] = s
	return domain.ReconcileResult{State: s, Applied: true}, nil
}
func (r *Reconciler) State(sku, warehouse string) (domain.ReconcileState, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s, ok := r.states[sku+"@"+warehouse]
	return s, ok
}
func (r *Reconciler) States() []domain.ReconcileState {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]domain.ReconcileState, 0, len(r.states))
	for _, s := range r.states {
		out = append(out, s)
	}
	return out
}
