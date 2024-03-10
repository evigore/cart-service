package cart

type addProductRequest struct {
	Count uint64 `json:"count"`
}

type getResponse struct {
	Items      []getItemResponse `json:"items"`
	TotalPrice uint64            `json:"total_price"`
}

type getItemResponse struct {
	Sku   int64  `json:"sku_id"`
	Name  string `json:"name"`
	Count uint64 `json:"count"`
	Price uint64 `json:"price"`
}
