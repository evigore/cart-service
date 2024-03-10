package cart

type GetCartResponse struct {
	Items      []GetCartItemResponse
	TotalPrice uint64
}
