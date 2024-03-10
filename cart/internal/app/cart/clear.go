package cart

import (
	"context"
	"fmt"
)

func (s CartService) Clear(ctx context.Context, userId int64) error {
	err := s.cartRepo.Clear(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed s.cartRepo.Clear: %w", err)
	}

	return nil
}
