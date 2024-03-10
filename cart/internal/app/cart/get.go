package cart

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"route256.ozon.ru/project/cart/internal/app/domain"
)

type GetCartItemResponse struct {
	Sku   int64
	Name  string
	Count uint64
	Price uint64
}

func (s CartService) Get(ctx context.Context, userId int64) (*GetCartResponse, error) {
	cart, err := s.cartRepo.Get(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("failed s.cartRepo.Get: %w", err)
	}

	if len(cart.Products) == 0 {
		return nil, fmt.Errorf("%w", domain.ErrNotFound)
	}

	// sort products by SKU
	slices.SortFunc(cart.Products, func(a, b domain.CartProduct) int {
		return cmp.Compare(a.Sku, b.Sku)
	})

	var totalPrice uint64
	items := make([]GetCartItemResponse, 0, len(cart.Products))
	for _, p := range cart.Products {
		product, err := s.productRepo.GetProductBySku(ctx, p.Sku)
		if err != nil {
			return nil, fmt.Errorf("failed s.productRepo.GetProductBySku: %w", err)
		}

		item := GetCartItemResponse{
			Sku:   p.Sku,
			Name:  product.Name,
			Count: p.Count,
			Price: product.Price,
		}

		totalPrice += uint64(p.Count) * product.Price
		items = append(items, item)
	}

	result := GetCartResponse{
		Items:      items,
		TotalPrice: totalPrice,
	}

	return &result, nil
}
