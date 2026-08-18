package serialize

import (
	"bytes"
	"encoding/json"
)

func Canonical(v any) ([]byte, error) {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	e.SetEscapeHTML(false)
	e.SetIndent("", "  ")
	if err := e.Encode(v); err != nil {
		return nil, err
	}
	return bytes.TrimSpace(b.Bytes()), nil
}
func Equal(a, b any) bool {
	left, err := Canonical(a)
	if err != nil {
		return false
	}
	right, err := Canonical(b)
	if err != nil {
		return false
	}
	return bytes.Equal(left, right)
}
func Must(v any) []byte {
	b, err := Canonical(v)
	if err != nil {
		panic(err)
	}
	return b
}
