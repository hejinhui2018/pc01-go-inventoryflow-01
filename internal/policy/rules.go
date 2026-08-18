package policy

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Rule struct {
	Name  string
	Check func(domain.InventoryEvent) error
}

func DefaultRules(l Limits) []Rule {
	return []Rule{{Name: "quantity", Check: func(e domain.InventoryEvent) error {
		q := e.Quantity
		if q < 0 {
			q = -q
		}
		return l.Check(q)
	}}, {Name: "version", Check: func(e domain.InventoryEvent) error {
		if e.Version <= 0 {
			return errors.New("version must be positive")
		}
		return nil
	}}}
}
func ValidateEvent(e domain.InventoryEvent, rules []Rule) error {
	for _, r := range rules {
		if err := r.Check(e); err != nil {
			return errors.New(r.Name + ": " + err.Error())
		}
	}
	return nil
}
func RequiresReview(e domain.InventoryEvent) bool {
	return e.Kind == domain.EventAdjustment || e.Quantity < 0
}
