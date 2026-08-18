package inventory

import (
	"errors"
	"github.com/hejinhui2018/pc01-go-inventoryflow-01/internal/domain"
	"sync"
)

type Catalog struct {
	mu         sync.RWMutex
	products   map[string]domain.Product
	warehouses map[string]domain.Warehouse
}

func NewCatalog() *Catalog {
	return &Catalog{products: map[string]domain.Product{}, warehouses: map[string]domain.Warehouse{}}
}
func (c *Catalog) AddProduct(p domain.Product) error {
	if err := p.Validate(); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.products[p.SKU]; ok {
		return errors.New("product already exists")
	}
	c.products[p.SKU] = p
	return nil
}
func (c *Catalog) AddWarehouse(w domain.Warehouse) error {
	if err := w.Validate(); err != nil {
		return err
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.warehouses[w.Code]; ok {
		return errors.New("warehouse already exists")
	}
	c.warehouses[w.Code] = w
	return nil
}
func (c *Catalog) Product(sku string) (domain.Product, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, ok := c.products[sku]
	return p, ok
}
func (c *Catalog) Warehouse(code string) (domain.Warehouse, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	w, ok := c.warehouses[code]
	return w, ok
}
func (c *Catalog) Products() []domain.Product {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make([]domain.Product, 0, len(c.products))
	for _, p := range c.products {
		out = append(out, p)
	}
	return out
}
