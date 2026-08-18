package domain

import (
	"errors"
	"strings"
	"time"
)

type OrderStatus string

const (
	OrderDraft     OrderStatus = "draft"
	OrderConfirmed OrderStatus = "confirmed"
	OrderPacked    OrderStatus = "packed"
	OrderShipped   OrderStatus = "shipped"
	OrderCancelled OrderStatus = "cancelled"
)

type OrderLine struct {
	SKU      string `json:"sku"`
	Quantity int64  `json:"quantity"`
}
type Order struct {
	ID        string      `json:"id"`
	Customer  string      `json:"customer"`
	Warehouse string      `json:"warehouse"`
	Lines     []OrderLine `json:"lines"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (o Order) Validate() error {
	if strings.TrimSpace(o.ID) == "" {
		return errors.New("order id is required")
	}
	if strings.TrimSpace(o.Customer) == "" {
		return errors.New("customer is required")
	}
	if strings.TrimSpace(o.Warehouse) == "" {
		return errors.New("warehouse is required")
	}
	if len(o.Lines) == 0 {
		return errors.New("order lines are required")
	}
	for _, line := range o.Lines {
		if line.SKU == "" || line.Quantity <= 0 {
			return errors.New("order line is invalid")
		}
	}
	return nil
}

func (o Order) CanTransition(to OrderStatus) bool {
	if o.Status == OrderDraft {
		return to == OrderConfirmed || to == OrderCancelled
	}
	if o.Status == OrderConfirmed {
		return to == OrderPacked || to == OrderCancelled
	}
	if o.Status == OrderPacked {
		return to == OrderShipped
	}
	return false
}

func (o Order) TotalUnits() int64 {
	var total int64
	for _, line := range o.Lines {
		total += line.Quantity
	}
	return total
}
