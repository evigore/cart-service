package repository

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/app/cart"
	"route256.ozon.ru/project/cart/internal/app/domain"
	"route256.ozon.ru/project/cart/internal/pkg/client/product"
	"route256.ozon.ru/project/cart/internal/pkg/utils/round-tripper/retry"
)

type ProductServiceRepository struct {
	client *product.Client
}

var _ cart.ProductRepository = (*ProductServiceRepository)(nil)

func NewProductServiceRepository(client *product.Client) *ProductServiceRepository {
	return &ProductServiceRepository{
		client,
	}
}

func (r ProductServiceRepository) GetProductBySku(ctx context.Context, sku int64) (*domain.Product, error) {
	r.client.Client.Transport = retry.New(3)

	res, err := r.client.GetProductBySku(ctx, sku)
	if err != nil {
		return nil, fmt.Errorf("failed r.client.GetProductBySku: %w", err)
	}

	if res.Code != product.CodeSuccess {
		if res.Code == product.CodeNotFound {
			return nil, domain.ErrNotFound
		}

		return nil, fmt.Errorf("unknown response status r.client.GetProductBySku: %d, %s", res.Code, res.Message)
	}

	product := domain.Product{
		Sku:   sku,
		Name:  res.Name,
		Price: res.Price,
	}

	return &product, nil
}
