package reconcile

import (
	"fmt"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Action struct {
	Version     int64
	EventID     string
	Description string
}
type Plan struct {
	Key          string
	Actions      []Action
	FinalVersion int64
}

func MakePlan(events []domain.InventoryEvent, applied int64) Plan {
	p := Plan{}
	for _, e := range Pending(events, applied) {
		if p.Key == "" {
			p.Key = e.SKU + "@" + e.Warehouse
		}
		p.Actions = append(p.Actions, Action{Version: e.Version, EventID: e.ID, Description: fmt.Sprintf("apply %s quantity %d", e.Kind, e.Quantity)})
		if e.Version > p.FinalVersion {
			p.FinalVersion = e.Version
		}
	}
	return p
}
func (p Plan) Empty() bool { return len(p.Actions) == 0 }
func (p Plan) Versions() []int64 {
	out := make([]int64, len(p.Actions))
	for i, a := range p.Actions {
		out[i] = a.Version
	}
	return out
}
