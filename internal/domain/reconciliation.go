package domain

import "time"

type ReconcileState struct {
	SKU            string    `json:"sku"`
	Warehouse      string    `json:"warehouse"`
	AppliedVersion int64     `json:"applied_version"`
	OnHand         int64     `json:"on_hand"`
	Reserved       int64     `json:"reserved"`
	UpdatedAt      time.Time `json:"updated_at"`
}
type ReconcileResult struct {
	State     ReconcileState `json:"state"`
	Applied   bool           `json:"applied"`
	Duplicate bool           `json:"duplicate"`
	Reason    string         `json:"reason,omitempty"`
}

func (s ReconcileState) Available() int64 { return s.OnHand - s.Reserved }
func (s ReconcileState) Healthy() bool {
	return s.OnHand >= 0 && s.Reserved >= 0 && s.Reserved <= s.OnHand
}
