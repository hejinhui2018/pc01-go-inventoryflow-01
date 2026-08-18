package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
)

type Availability struct {
	OnHand   int64
	Reserved int64
}

func (a Availability) Available() int64 { return a.OnHand - a.Reserved }
func (a Availability) CanFulfill(quantity int64) bool {
	return quantity > 0 && a.Available() >= quantity
}
func (a Availability) Validate() error {
	if a.OnHand < 0 {
		return errors.New("on-hand cannot be negative")
	}
	if a.Reserved < 0 {
		return errors.New("reserved cannot be negative")
	}
	if a.Reserved > a.OnHand {
		return errors.New("reserved exceeds on-hand")
	}
	return nil
}
func AvailabilityFrom(s domain.ReconcileState) Availability {
	return Availability{OnHand: s.OnHand, Reserved: s.Reserved}
}
func (b *ReservationBook) Availability(sku string, onHand int64) Availability {
	return Availability{OnHand: onHand, Reserved: b.Reserved(sku, "", nil)}
}
