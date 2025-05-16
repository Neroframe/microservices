package inmemory

import (
	"log"
	"sync"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

var _ domain.ProductMemoryCache = (*ProductCache)(nil)

type ProductCache struct {
	products map[string]*domain.Product
	m        sync.RWMutex
}

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
	log.Printf("[InMemory] Set product id=%s", product.ID)
}

func (c *ProductCache) SetMany(products []*domain.Product) {
	c.m.Lock()
	defer c.m.Unlock()

	for _, product := range products {
		c.products[product.ID] = product
		log.Printf("[InMemory] SetMany product id=%s", product.ID)
	}
	log.Printf("[InMemory] SetMany done: total=%d", len(products))
}

func (c *ProductCache) Get(productID string) (*domain.Product, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	product, ok := c.products[productID]
	if ok {
		log.Printf("[InMemory] HIT for id=%s", productID)
	} else {
		log.Printf("[InMemory] MISS for id=%s", productID)
	}
	return product, ok
}

func (c *ProductCache) Delete(productID string) {
	c.m.Lock()
	defer c.m.Unlock()

	delete(c.products, productID)
	log.Printf("[InMemory] Deleted product id=%s", productID)
}

func (c *ProductCache) GetList() ([]*domain.Product, bool) {
	c.m.RLock()
	defer c.m.RUnlock()

	if len(c.products) == 0 {
		log.Printf("[InMemory] GetList: empty cache")
		return nil, false
	}

	list := make([]*domain.Product, 0, len(c.products))
	for _, product := range c.products {
		list = append(list, product)
	}

	log.Printf("[InMemory] GetList: returned %d products", len(list))
	return list, true
}
