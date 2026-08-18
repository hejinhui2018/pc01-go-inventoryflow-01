package ids

import (
	"crypto/rand"
	"encoding/hex"
)

func New(prefix string) string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return prefix + "-fallback"
	}
	return prefix + "-" + hex.EncodeToString(b)
}
