package cart

type parsedRequestGet struct {
	UserId int64
}

type parsedRequestAddProduct struct {
	UserId int64
	Sku    int64
	Count  uint64
}

type parsedRequestDeleteProduct struct {
	UserId int64
	Sku    int64
}

type parsedRequestClear struct {
	UserId int64
}
