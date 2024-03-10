package cart

import (
	"context"

	"route256.ozon.ru/project/cart/internal/app/domain"
)

type CartRepository interface {
	Get(ctx context.Context, userId int64) (*domain.Cart, error)
	AddProduct(ctx context.Context, userId int64, sku int64, count uint64) error
	DeleteProduct(ctx context.Context, userId int64, sku int64) error
	Clear(ctx context.Context, userId int64) error
}

type ProductRepository interface {
	GetProductBySku(ctx context.Context, sku int64) (*domain.Product, error)
}

type CartService struct {
	cartRepo    CartRepository
	productRepo ProductRepository
}

func New(
	cartRepo CartRepository,
	productRepo ProductRepository,
) *CartService {
	return &CartService{
		cartRepo,
		productRepo,
	}
}
