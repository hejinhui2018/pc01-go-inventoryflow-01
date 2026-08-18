package workflow

import (
	"encoding/json"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"io"
)

func ImportEvents(r io.Reader) ([]domain.InventoryEvent, error) {
	var events []domain.InventoryEvent
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&events); err != nil {
		return nil, err
	}
	for _, e := range events {
		if err := e.Validate(); err != nil {
			return nil, err
		}
	}
	return events, nil
}
func ExportEvents(w io.Writer, events []domain.InventoryEvent) error {
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(events)
}
func (s *Service) Import(r io.Reader) (int, error) {
	events, err := ImportEvents(r)
	if err != nil {
		return 0, err
	}
	n := 0
	for _, e := range events {
		if _, err := s.Reconcile(e); err != nil {
			return n, err
		}
		n++
	}
	return n, nil
}
