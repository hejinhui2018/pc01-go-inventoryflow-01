package domain

import (
	"errors"
	"time"
)

type ReservationStatus string

const (
	ReservationOpen      ReservationStatus = "open"
	ReservationCommitted ReservationStatus = "committed"
	ReservationReleased  ReservationStatus = "released"
)

type Reservation struct {
	ID        string            `json:"id"`
	OrderID   string            `json:"order_id"`
	SKU       string            `json:"sku"`
	Quantity  int64             `json:"quantity"`
	Status    ReservationStatus `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
}

func (r Reservation) Validate() error {
	if r.ID == "" || r.OrderID == "" || r.SKU == "" {
		return errors.New("reservation identity is required")
	}
	if r.Quantity <= 0 {
		return errors.New("reservation quantity must be positive")
	}
	if r.Status == "" {
		return errors.New("reservation status is required")
	}
	return nil
}

func (r Reservation) CanRelease() bool { return r.Status == ReservationOpen }
func (r Reservation) CanCommit() bool  { return r.Status == ReservationOpen }
