package workflow

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"time"
)

type OrderBook struct{ orders map[string]domain.Order }

func NewOrderBook() *OrderBook { return &OrderBook{orders: map[string]domain.Order{}} }
func (b *OrderBook) Create(o domain.Order) error {
	if err := o.Validate(); err != nil {
		return err
	}
	if _, ok := b.orders[o.ID]; ok {
		return errors.New("order already exists")
	}
	if o.Status == "" {
		o.Status = domain.OrderDraft
	}
	if o.CreatedAt.IsZero() {
		o.CreatedAt = time.Now().UTC()
	}
	o.UpdatedAt = o.CreatedAt
	b.orders[o.ID] = o
	return nil
}
func (b *OrderBook) Transition(id string, to domain.OrderStatus) (domain.Order, error) {
	o, ok := b.orders[id]
	if !ok {
		return domain.Order{}, errors.New("order not found")
	}
	if !o.CanTransition(to) {
		return o, errors.New("invalid order transition")
	}
	o.Status = to
	o.UpdatedAt = time.Now().UTC()
	b.orders[id] = o
	return o, nil
}
func (b *OrderBook) Get(id string) (domain.Order, bool) { o, ok := b.orders[id]; return o, ok }
func (b *OrderBook) All() []domain.Order {
	out := make([]domain.Order, 0, len(b.orders))
	for _, o := range b.orders {
		out = append(out, o)
	}
	return out
}
