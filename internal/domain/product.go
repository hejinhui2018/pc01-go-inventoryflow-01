package domain

import "errors"

type Product struct {
	SKU    string `json:"sku"`
	Name   string `json:"name"`
	Unit   string `json:"unit"`
	Active bool   `json:"active"`
}

func (p Product) Validate() error {
	if p.SKU == "" {
		return errors.New("sku is required")
	}
	if p.Name == "" {
		return errors.New("name is required")
	}
	if p.Unit == "" {
		return errors.New("unit is required")
	}
	return nil
}
