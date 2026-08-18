package serialize

import (
	"encoding/json"
	"io"
)

func Read(r io.Reader, out any) error {
	d := json.NewDecoder(r)
	d.DisallowUnknownFields()
	return d.Decode(out)
}
func Write(w io.Writer, v any) error {
	e := json.NewEncoder(w)
	e.SetEscapeHTML(false)
	return e.Encode(v)
}
