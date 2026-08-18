package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Movement struct {
	Before int64
	Delta  int64
	After  int64
	Event  domain.InventoryEvent
}

func BuildMovement(before int64, e domain.InventoryEvent) (Movement, error) {
	if err := e.Validate(); err != nil {
		return Movement{}, err
	}
	after := before + e.Delta()
	if after < 0 {
		return Movement{}, errors.New("movement would make stock negative")
	}
	return Movement{Before: before, Delta: e.Delta(), After: after, Event: e}, nil
}
func (m Movement) Reversible() bool { return m.Before >= 0 && m.After >= 0 }
func (m Movement) Reverse() Movement {
	return Movement{Before: m.After, Delta: -m.Delta, After: m.Before, Event: m.Event}
}
func ApplyMovements(start int64, events []domain.InventoryEvent) (int64, error) {
	total := start
	for _, e := range events {
		m, err := BuildMovement(total, e)
		if err != nil {
			return total, err
		}
		total = m.After
	}
	return total, nil
}
