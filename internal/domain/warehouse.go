package domain

import "errors"

type Warehouse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Region  string `json:"region"`
	Enabled bool   `json:"enabled"`
}

func (w Warehouse) Validate() error {
	if w.Code == "" {
		return errors.New("warehouse code is required")
	}
	if w.Name == "" {
		return errors.New("warehouse name is required")
	}
	if w.Region == "" {
		return errors.New("warehouse region is required")
	}
	return nil
}
