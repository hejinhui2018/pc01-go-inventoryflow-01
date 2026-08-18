package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sync"
)

type ReservationBook struct {
	mu    sync.RWMutex
	items map[string]domain.Reservation
}

func NewReservationBook() *ReservationBook {
	return &ReservationBook{items: make(map[string]domain.Reservation)}
}
func (b *ReservationBook) Create(r domain.Reservation) error {
	if err := r.Validate(); err != nil {
		return err
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.items[r.ID]; ok {
		return errors.New("reservation already exists")
	}
	b.items[r.ID] = r
	return nil
}
func (b *ReservationBook) Change(id string, status domain.ReservationStatus) error {
	b.mu.Lock()
	defer b.mu.Unlock()
	r, ok := b.items[id]
	if !ok {
		return errors.New("reservation not found")
	}
	if status == domain.ReservationCommitted && !r.CanCommit() {
		return errors.New("reservation cannot commit")
	}
	if status == domain.ReservationReleased && !r.CanRelease() {
		return errors.New("reservation cannot release")
	}
	r.Status = status
	b.items[id] = r
	return nil
}
func (b *ReservationBook) Get(id string) (domain.Reservation, bool) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	r, ok := b.items[id]
	return r, ok
}
func (b *ReservationBook) All() []domain.Reservation {
	b.mu.RLock()
	defer b.mu.RUnlock()
	out := make([]domain.Reservation, 0, len(b.items))
	for _, r := range b.items {
		out = append(out, r)
	}
	return out
}
func (b *ReservationBook) Reserved(sku, warehouse string, _ func(domain.Reservation) bool) int64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	var total int64
	for _, r := range b.items {
		if r.SKU == sku && r.Status == domain.ReservationOpen {
			total += r.Quantity
		}
	}
	return total
}
