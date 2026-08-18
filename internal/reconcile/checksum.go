package reconcile

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"strconv"
)

func EventChecksum(e domain.InventoryEvent) string {
	h := sha256.New()
	h.Write([]byte(e.ID))
	h.Write([]byte(e.SKU))
	h.Write([]byte(e.Warehouse))
	h.Write([]byte(strconv.FormatInt(e.Version, 10)))
	h.Write([]byte(strconv.FormatInt(e.Quantity, 10)))
	return hex.EncodeToString(h.Sum(nil))
}
func StateChecksum(s domain.ReconcileState) string {
	h := sha256.New()
	h.Write([]byte(s.SKU))
	h.Write([]byte(s.Warehouse))
	h.Write([]byte(strconv.FormatInt(s.AppliedVersion, 10)))
	h.Write([]byte(strconv.FormatInt(s.OnHand, 10)))
	return hex.EncodeToString(h.Sum(nil))
}
