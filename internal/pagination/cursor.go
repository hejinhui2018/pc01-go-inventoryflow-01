package pagination

import (
	"encoding/base64"
	"strconv"
)

func Encode(offset int) string {
	return base64.RawURLEncoding.EncodeToString([]byte(strconv.Itoa(offset)))
}
func Decode(value string) (int, bool) {
	if value == "" {
		return 0, true
	}
	b, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return 0, false
	}
	v, err := strconv.Atoi(string(b))
	return v, err == nil && v >= 0
}
func Next[T any](items []T, offset, limit int) (Page[T], string) {
	p := Build(items, offset, limit)
	next := p.Offset + len(p.Items)
	if next >= p.Total {
		return p, ""
	}
	return p, Encode(next)
}
