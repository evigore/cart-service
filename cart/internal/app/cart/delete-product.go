package cart

import (
	"context"
	"fmt"
)

func (s CartService) DeleteProduct(ctx context.Context, userId int64, sku int64) error {
	err := s.cartRepo.DeleteProduct(ctx, userId, sku)
	if err != nil {
		return fmt.Errorf("failed s.cartRepo.DeleteProduct: %w", err)
	}

	return nil
}
