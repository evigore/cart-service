package domain

type CartProduct struct {
	Sku   int64
	Count uint64
}

type Cart struct {
	UserId   int64
	Products []CartProduct
}
