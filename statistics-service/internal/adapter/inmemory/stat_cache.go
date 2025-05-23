package cache

import (
	"log"
	"sync"

	"github.com/Neroframe/ecommerce-platform/statistics-service/internal/domain"
)

type InMemoryEventCache struct {
	mu    sync.RWMutex
	items map[string]*domain.Event // key = EntityID
	order []string                 // keep insertion order
}

func NewInMemoryEventCache() *InMemoryEventCache {
	return &InMemoryEventCache{
		items: make(map[string]*domain.Event),
		order: make([]string, 0),
	}
}

func (c *InMemoryEventCache) Get(id string) (*domain.Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	evt, ok := c.items[id]
	return evt, ok
}

func (c *InMemoryEventCache) Set(evt *domain.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[evt.EntityID]; !exists {
		c.order = append(c.order, evt.EntityID)
	}
	c.items[evt.EntityID] = evt
	log.Printf("[InMemory] Set event id=%s type=%s", evt.EntityID, evt.EventType)
}

func (c *InMemoryEventCache) SetMany(evts []*domain.Event) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = make(map[string]*domain.Event, len(evts))
	c.order = make([]string, 0, len(evts))

	for _, evt := range evts {
		c.items[evt.EntityID] = evt
		c.order = append(c.order, evt.EntityID)
	}
	log.Printf("[InMemory] SetMany events count=%d", len(evts))
}

func (c *InMemoryEventCache) Delete(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, id)
	for i, eid := range c.order {
		if eid == id {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	log.Printf("[InMemory] Deleted event id=%s", id)
}

func (c *InMemoryEventCache) GetList() ([]*domain.Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.items) == 0 {
		return nil, false
	}
	out := make([]*domain.Event, 0, len(c.order))
	for _, id := range c.order {
		if evt, ok := c.items[id]; ok {
			out = append(out, evt)
		}
	}
	return out, true
}
