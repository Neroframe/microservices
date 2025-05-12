package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

// productUsecase wraps repository and Redis client for caching
// and implements domain.ProductUsecase
type productUsecase struct {
	productRepo domain.ProductRepository
	cache       *redis.Client
	cacheTTL    time.Duration
}

// NewProductUsecase constructs a usecase with caching support
func NewProductUsecase(repo domain.ProductRepository, cache *redis.Client, cacheTTL time.Duration) domain.ProductUsecase {
	return &productUsecase{productRepo: repo, cache: cache, cacheTTL: cacheTTL}
}

// Create persists a new product, caches its detail, and invalidates the list
func (u *productUsecase) Create(ctx context.Context, p *domain.Product) error {
	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if len(p.Name) > 36 {
		return errors.New("product name cannot exceed 36 characters")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Stock < 0 {
		return errors.New("stock value cannot be negative")
	}

	p.NormalizeName()

	// save to DB
	err := u.productRepo.Create(ctx, p)
	if err != nil {
		return err
	}

	// cache detail
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("cache marshal detail error: %v\n", err)
	} else {
		if err := u.cache.Set(ctx, detailKey(p.ID), data, u.cacheTTL).Err(); err != nil {
			fmt.Printf("cache set detail error: %v\n", err)
		}
	}

	// invalidate list cache
	if err := u.cache.Del(ctx, listKey()).Err(); err != nil {
		fmt.Printf("cache delete list error: %v\n", err)
	}

	return nil
}

// List returns all products, attempting cache first
func (u *productUsecase) List(ctx context.Context) ([]*domain.Product, error) {
	// try cache
	data, err := u.cache.Get(ctx, listKey()).Result()
	if err == nil {
		var prods []*domain.Product
		if err := json.Unmarshal([]byte(data), &prods); err == nil {
			return prods, nil
		}
		fmt.Printf("cache unmarshal list error: %v\n", err)
	}

	// cache miss: fetch from DB
	prods, err := u.productRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	// cache list
	b, err := json.Marshal(prods)
	if err != nil {
		fmt.Printf("cache marshal list error: %v\n", err)
	} else {
		if err := u.cache.Set(ctx, listKey(), b, u.cacheTTL).Err(); err != nil {
			fmt.Printf("cache set list error: %v\n", err)
		}
	}

	return prods, nil
}

// GetByID returns a single product, attempting cache first
func (u *productUsecase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	// try cache
	data, err := u.cache.Get(ctx, detailKey(id)).Result()
	if err == nil {
		var p domain.Product
		if err := json.Unmarshal([]byte(data), &p); err == nil {
			return &p, nil
		}
		fmt.Printf("cache unmarshal detail error: %v\n", err)
	}

	// cache miss: fetch from DB
	p, err := u.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// cache detail
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("cache marshal detail error: %v\n", err)
	} else {
		if err := u.cache.Set(ctx, detailKey(id), b, u.cacheTTL).Err(); err != nil {
			fmt.Printf("cache set detail error: %v\n", err)
		}
	}

	return p, nil
}

// RefreshProductsCache reloads all products from DB into cache (list and details)
func (u *productUsecase) RefreshProductsCache(ctx context.Context) error {
	// fetch all products
	prods, err := u.productRepo.List(ctx)
	if err != nil {
		return err
	}

	// cache list
	b, err := json.Marshal(prods)
	if err != nil {
		fmt.Printf("cache marshal list error: %v\n", err)
	} else {
		if err := u.cache.Set(ctx, listKey(), b, u.cacheTTL).Err(); err != nil {
			fmt.Printf("cache set list error: %v\n", err)
		}
	}

	// cache each detail
	for _, p := range prods {
		bp, err := json.Marshal(p)
		if err != nil {
			fmt.Printf("cache marshal detail error for id %s: %v\n", p.ID, err)
			continue
		}
		if err := u.cache.Set(ctx, detailKey(p.ID), bp, u.cacheTTL).Err(); err != nil {
			fmt.Printf("cache set detail error for id %s: %v\n", p.ID, err)
		}
	}

	return nil
}

// listKey returns the Redis key for the product list
func listKey() string {
	return "inventory:products:list"
}

// detailKey returns the Redis key for an individual product
func detailKey(id string) string {
	return fmt.Sprintf("inventory:products:%s", id)
}

func (u *productUsecase) Update(ctx context.Context, p *domain.Product) error {
	if p.ID == "" {
		return errors.New("product ID cannot be empty")
	}

	if p.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if len(p.Name) > 36 {
		return errors.New("product name cannot exceed 36 characters")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if p.Stock < 0 {
		return errors.New("stock value cannot be negative")
	}

	p.NormalizeName()

	return u.productRepo.Update(ctx, p)
}

func (u *productUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}
	return u.productRepo.Delete(ctx, id)
}
