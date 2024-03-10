package product

type GetProductRequest struct {
	Token string `json:"token"`
	Sku   int64  `json:"sku"`
}

type GetProductResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`

	Name  string `json:"name"`
	Price uint64 `json:"price"`
}
