package reconcile

import (
	"context"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Engine struct{ reconciler *Reconciler }

func NewEngine(r *Reconciler) *Engine { return &Engine{reconciler: r} }
func (e *Engine) Run(ctx context.Context, events []domain.InventoryEvent, reserved func(domain.InventoryEvent) int64) ([]domain.ReconcileResult, error) {
	out := make([]domain.ReconcileResult, 0, len(events))
	for _, item := range SortEvents(events) {
		select {
		case <-ctx.Done():
			return out, ctx.Err()
		default:
		}
		result, err := e.reconciler.Apply(item, reserved(item))
		if err != nil {
			return out, err
		}
		out = append(out, result)
	}
	return out, nil
}
func (e *Engine) DryRun(events []domain.InventoryEvent, start domain.ReconcileState) domain.ReconcileState {
	state := start
	for _, item := range SortEvents(events) {
		state.OnHand += item.Delta()
		if item.Version > state.AppliedVersion {
			state.AppliedVersion = item.Version
		}
	}
	return state
}
func (e *Engine) Reconciler() *Reconciler { return e.reconciler }
