package cart

import (
	"context"
	"fmt"
)

func (s CartService) AddProduct(ctx context.Context, userId int64, sku int64, count uint64) error {
	_, err := s.productRepo.GetProductBySku(ctx, sku)
	if err != nil {
		return fmt.Errorf("failed s.productRepo.GetProductBySku: %w", err)
	}

	return s.cartRepo.AddProduct(ctx, userId, sku, count)
}
