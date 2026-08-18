package domain

import (
	"errors"
	"time"
)

type TransferStatus string

const (
	TransferRequested TransferStatus = "requested"
	TransferInTransit TransferStatus = "in_transit"
	TransferReceived  TransferStatus = "received"
	TransferRejected  TransferStatus = "rejected"
)

type Transfer struct {
	ID          string         `json:"id"`
	SKU         string         `json:"sku"`
	From        string         `json:"from"`
	To          string         `json:"to"`
	Quantity    int64          `json:"quantity"`
	Status      TransferStatus `json:"status"`
	RequestedAt time.Time      `json:"requested_at"`
}

func (t Transfer) Validate() error {
	if t.ID == "" || t.SKU == "" || t.From == "" || t.To == "" {
		return errors.New("transfer identity is required")
	}
	if t.From == t.To {
		return errors.New("transfer endpoints must differ")
	}
	if t.Quantity <= 0 {
		return errors.New("transfer quantity must be positive")
	}
	return nil
}
func (t Transfer) Start() (Transfer, error) {
	if t.Status != TransferRequested {
		return t, errors.New("transfer is not requested")
	}
	t.Status = TransferInTransit
	return t, nil
}
func (t Transfer) Receive() (Transfer, error) {
	if t.Status != TransferInTransit {
		return t, errors.New("transfer is not in transit")
	}
	t.Status = TransferReceived
	return t, nil
}
func (t Transfer) Reject() (Transfer, error) {
	if t.Status == TransferReceived {
		return t, errors.New("received transfer cannot be rejected")
	}
	t.Status = TransferRejected
	return t, nil
}
