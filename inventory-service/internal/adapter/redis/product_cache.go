package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
	"github.com/Neroframe/ecommerce-platform/inventory-service/pkg/redis"
	goredis "github.com/redis/go-redis/v9"
)

var _ domain.ProductRedisCache = (*ProductCache)(nil)

const keyPrefix = "product:%s"
const productListKey = "product:list"

type ProductCache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewProductCache(client *redis.Client, ttl time.Duration) *ProductCache {
	return &ProductCache{
		client: client,
		ttl:    ttl,
	}
}

func (c *ProductCache) Set(ctx context.Context, product *domain.Product) error {
	data, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("failed to marshal product Set: %w", err)
	}

	key := c.key(product.ID)
	err = c.client.Unwrap().Set(ctx, key, data, c.ttl).Err()
	if err != nil {
		return fmt.Errorf("redis Set error: %w", err)
	}

	log.Printf("[Redis] Set product key=%s ttl=%s", key, c.ttl)
	return nil
}

func (c *ProductCache) SetMany(ctx context.Context, products []*domain.Product) error {
	pipe := c.client.Unwrap().Pipeline()
	for _, product := range products {
		data, err := json.Marshal(product)
		if err != nil {
			return fmt.Errorf("failed to marshal product SetMany: %w", err)
		}
		pipe.Set(ctx, c.key(product.ID), data, c.ttl)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to set many products: %w", err)
	}

	log.Printf("[Redis] SetMany committed %d products", len(products))
	return nil
}

func (c *ProductCache) Get(ctx context.Context, productID string) (*domain.Product, error) {
	key := c.key(productID)
	data, err := c.client.Unwrap().Get(ctx, c.key(productID)).Bytes()
	if err != nil {
		if err == goredis.Nil {
			log.Printf("[Redis] MISS for key=%s", key)
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	log.Printf("[Redis] HIT for key=%s", key)

	var product domain.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal product: %w", err)
	}

	return &product, nil
}

func (c *ProductCache) Delete(ctx context.Context, productID string) error {
	return c.client.Unwrap().Del(ctx, c.key(productID)).Err()
}

func (c *ProductCache) SetList(ctx context.Context, products []*domain.Product) error {
	data, err := json.Marshal(products)
	if err != nil {
		return fmt.Errorf("marshal product list: %w", err)
	}

	err = c.client.Unwrap().Set(ctx, productListKey, data, c.ttl).Err()
	if err != nil {
		return fmt.Errorf("redis SetList error: %w", err)
	}

	log.Printf("[Redis] Set product list key=%s count=%d ttl=%s", productListKey, len(products), c.ttl)
	return nil
}

func (c *ProductCache) GetList(ctx context.Context) ([]*domain.Product, error) {
	data, err := c.client.Unwrap().Get(ctx, productListKey).Bytes()
	if err != nil {
		if err == goredis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("get product list: %w", err)
	}

	var products []*domain.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, fmt.Errorf("unmarshal product list: %w", err)
	}

	return products, nil
}

func (c *ProductCache) key(id string) string {
	return fmt.Sprintf(keyPrefix, id)
}
