package repository

import (
	"context"
	"sync"

	"route256.ozon.ru/project/cart/internal/app/cart"
	"route256.ozon.ru/project/cart/internal/app/domain"
)

type CartInMemoryRepository struct {
	mu sync.RWMutex
	// map[userId]map[sku]count
	storage map[int64]map[int64]uint64
}

var _ cart.CartRepository = (*CartInMemoryRepository)(nil)

func NewCartInMemoryRepository() *CartInMemoryRepository {
	return &CartInMemoryRepository{
		storage: make(map[int64]map[int64]uint64, 0),
	}
}

func (r *CartInMemoryRepository) Get(ctx context.Context, userId int64) (*domain.Cart, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	skus, ok := r.storage[userId]
	if !ok {
		return &domain.Cart{UserId: userId}, nil
	}

	products := make([]domain.CartProduct, 0, len(skus))
	for sku, count := range skus {
		product := domain.CartProduct{
			Sku:   sku,
			Count: count,
		}

		products = append(products, product)
	}

	return &domain.Cart{UserId: userId, Products: products}, nil
}

func (r *CartInMemoryRepository) AddProduct(ctx context.Context, userId int64, sku int64, count uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.storage[userId]
	if !ok {
		r.storage[userId] = make(map[int64]uint64, 1)
	}

	r.storage[userId][sku] += count
	return nil
}

func (r *CartInMemoryRepository) DeleteProduct(ctx context.Context, userId int64, sku int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.storage[userId]
	if ok {
		delete(r.storage[userId], sku)
	}

	return nil
}

func (r *CartInMemoryRepository) Clear(ctx context.Context, userId int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.storage[userId]
	if ok {
		clear(r.storage[userId])
	}

	return nil
}
