package cart

import "route256.ozon.ru/project/cart/internal/app/cart"

func getResponseConverter(cart *cart.GetCartResponse) *getResponse {
	items := make([]getItemResponse, 0, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, getItemResponse{
			Sku:   item.Sku,
			Name:  item.Name,
			Count: item.Count,
			Price: item.Price,
		})
	}

	return &getResponse{
		Items:      items,
		TotalPrice: cart.TotalPrice,
	}
}
