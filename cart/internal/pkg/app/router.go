package app

import (
	"net/http"

	cartService "route256.ozon.ru/project/cart/internal/app/cart"
	"route256.ozon.ru/project/cart/internal/config"
	"route256.ozon.ru/project/cart/internal/pkg/client/product"
	"route256.ozon.ru/project/cart/internal/pkg/controller/cart"
	"route256.ozon.ru/project/cart/internal/pkg/middleware"
	"route256.ozon.ru/project/cart/internal/pkg/repository"
)

func setupRouter(mux *http.ServeMux, config *config.Config) http.Handler {
	cartInMemoryRepository := repository.NewCartInMemoryRepository()
	productServiceRepository := repository.NewProductServiceRepository(product.New(config.ProductServiceUrl, config.ProductServiceToken))

	cartController := cart.New(
		*cartService.New(cartInMemoryRepository, productServiceRepository),
	)

	mux.HandleFunc("GET /user/{user_id}/cart", cartController.Get)
	mux.HandleFunc("POST /user/{user_id}/cart/{sku...}", cartController.AddProduct)
	mux.HandleFunc("DELETE /user/{user_id}/cart/{sku...}", cartController.DeleteProduct)
	mux.HandleFunc("DELETE /user/{user_id}/cart", cartController.Clear)

	handler := middleware.LoggerMiddleware(mux)
	handler = middleware.ResponseMiddleware(handler)

	return handler
}
