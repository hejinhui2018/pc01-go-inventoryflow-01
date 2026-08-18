package store

import (
	"bytes"
	"encoding/json"
)

func Encode(value any) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}
func Decode(data []byte, out any) error {
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	return dec.Decode(out)
}
func Clone[T any](in T) (T, error) {
	var out T
	b, err := Encode(in)
	if err != nil {
		return out, err
	}
	err = Decode(b, &out)
	return out, err
}
