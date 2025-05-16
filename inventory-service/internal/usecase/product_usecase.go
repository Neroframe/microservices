package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/Neroframe/ecommerce-platform/inventory-service/internal/domain"
)

// productUsecase wraps repository and Redis client for caching
// and implements domain.ProductUsecase
type productUsecase struct {
	productRepo   domain.ProductRepository
	publisher     domain.InventoryEventPublisher
	inMemoryCache domain.ProductMemoryCache
	redisCache    domain.ProductRedisCache
}

// NewProductUsecase constructs a usecase with caching support
func NewProductUsecase(repo domain.ProductRepository, pub domain.InventoryEventPublisher, inmemory domain.ProductMemoryCache, redis domain.ProductRedisCache) domain.ProductUsecase {
	return &productUsecase{
		productRepo:   repo,
		publisher:     pub,
		inMemoryCache: inmemory,
		redisCache:    redis,
	}
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

	// set inmemory cache
	u.inMemoryCache.Set(p)

	// set redis cache
	if err := u.redisCache.Set(ctx, p); err != nil {
		return fmt.Errorf("redisCache.Set: %w", err)
	}

	// NATS publish
	event := domain.ProductCreatedEvent{
		ID:         p.ID,
		Name:       p.Name,
		Price:      int(p.Price),
		CategoryID: p.Category,
	}
	if err := u.publisher.PublishProductCreated(ctx, event); err != nil {
		return fmt.Errorf("publisher.PublishProductCreated: %w", err)
	}

	return nil
}

// List returns all products, attempting cache first
func (u *productUsecase) List(ctx context.Context) ([]*domain.Product, error) {
	// inmemory cache
	if products, ok := u.inMemoryCache.GetList(); ok {
		return products, nil
	}

	// redis cache
	products, err := u.redisCache.GetList(ctx)
	if err == nil && products != nil {
		u.inMemoryCache.SetMany(products) // warm memory
		return products, nil
	}

	// mongoDB
	products, err = u.productRepo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("productRepo.GetAll: %w", err)
	}

	// update cache
	u.inMemoryCache.SetMany(products)
	_ = u.redisCache.SetList(ctx, products)

	return products, nil
}

func (u *productUsecase) GetByID(ctx context.Context, id string) (*domain.Product, error) {
	// try inmemory
	if product, ok := u.inMemoryCache.Get(id); ok {
		return product, nil
	}

	// try Redis
	product, err := u.redisCache.Get(ctx, id)
	if err == nil && product != nil {
		u.inMemoryCache.Set(product) // warm in-memory
		return product, nil
	}

	//  DB
	product, err = u.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("productRepo.GetByID: %w", err)
	}
	if product == nil {
		return nil, nil // not found
	}

	// Warm both caches
	u.inMemoryCache.Set(product)
	_ = u.redisCache.Set(ctx, product)

	return product, nil
}

func (u *productUsecase) RefreshProductsCache(ctx context.Context) error {
	// 1. Load all products from DB
	products, err := u.productRepo.List(ctx)
	if err != nil {
		return fmt.Errorf("productRepo.List: %w", err)
	}

	// 2. Refresh in-memory cache
	u.inMemoryCache.SetMany(products)

	// 3. Refresh Redis cache
	if err := u.redisCache.SetList(ctx, products); err != nil {
		return fmt.Errorf("redisCache.SetList: %w", err)
	}

	return nil
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

	if err := u.productRepo.Update(ctx, p); err != nil {
		return err
	}

	// caches
	u.inMemoryCache.Set(p)
	_ = u.redisCache.Set(ctx, p)

	// publish
	evt := domain.ProductUpdatedEvent{
		ID:         p.ID,
		Name:       p.Name,
		Price:      int(p.Price),
		CategoryID: p.Category,
	}
	if err := u.publisher.PublishProductUpdated(ctx, evt); err != nil {
		return fmt.Errorf("publisher.PublishProductUpdated: %w", err)
	}

	return nil
}

func (u *productUsecase) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("product ID cannot be empty")
	}

	if err := u.productRepo.Delete(ctx, id); err != nil {
		return err
	}

	// invalidate caches
	u.inMemoryCache.Delete(id)
	err := u.redisCache.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("redisCache.Delete error: %w", err)
	}

	// publish
	if err := u.publisher.PublishProductDeleted(ctx, domain.ProductDeletedEvent{ID: id}); err != nil {
		return fmt.Errorf("publisher.PublishProductDeleted: %w", err)
	}

	return nil
}
