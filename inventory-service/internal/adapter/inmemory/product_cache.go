package inmemory

import (
	"sync"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

type ProductCache struct {
	products map[string]*domain.Product
	m        sync.RWMutex
}

var _ domain.ProductMemoryCache = (*ProductCache)(nil)

func NewProductCache() *ProductCache {
	return &ProductCache{
		products: make(map[string]*domain.Product),
		m:        sync.RWMutex{},
	}
}

func (c *ProductCache) Set(product *domain.Product) {
	c.m.Lock()
	defer c.m.Unlock()

	c.products[product.ID] = product
}

func (c *ProductCache) SetMany(products []*domain.Product) {
	c.m.Lock()
	defer c.m.Unlock()

	for _, product := range products {
		c.products[product.ID] = product
	}
}

func (c *ProductCache) Get(productID string) (*domain.Product, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	product, ok := c.products[productID]
	return product, ok
}

func (c *ProductCache) Delete(productID string) {
	c.m.Lock()
	defer c.m.Unlock()

	delete(c.products, productID)
}

func (c *ProductCache) GetList() ([]*domain.Product, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	if len(c.products) == 0 {
		return nil, false
	}

	list := make([]*domain.Product, 0, len(c.products))
	for _, product := range c.products {
		list = append(list, product)
	}
	return list, true
}
