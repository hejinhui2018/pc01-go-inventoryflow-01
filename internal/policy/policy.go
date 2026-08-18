package policy

import "errors"

type Limits struct {
	MaxReservation int64
	AllowNegative  bool
}

func DefaultLimits() Limits { return Limits{MaxReservation: 100000, AllowNegative: false} }
func (l Limits) Check(quantity int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be positive")
	}
	if quantity > l.MaxReservation {
		return errors.New("quantity exceeds policy limit")
	}
	return nil
}
