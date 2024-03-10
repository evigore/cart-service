package cart

import (
	"encoding/json"
	"errors"
	"net/http"

	"route256.ozon.ru/project/cart/internal/app/cart"
	"route256.ozon.ru/project/cart/internal/app/domain"
)

type CartController struct {
	service cart.CartService
}

func New(
	service cart.CartService,
) *CartController {
	return &CartController{
		service,
	}
}

func (c CartController) Get(w http.ResponseWriter, r *http.Request) {
	req, err := parseGetRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart, err := c.service.Get(r.Context(), req.UserId)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}

		http.Error(w, err.Error(), statusCode)
		return
	}

	bytes, err := json.Marshal(getResponseConverter(cart))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}

func (c CartController) Clear(w http.ResponseWriter, r *http.Request) {
	req, err := parseClearRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.Clear(r.Context(), req.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c CartController) AddProduct(w http.ResponseWriter, r *http.Request) {
	req, err := parseAddProductRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.AddProduct(r.Context(), req.UserId, req.Sku, req.Count)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, domain.ErrNotFound) {
			statusCode = http.StatusNotFound
		}

		http.Error(w, err.Error(), statusCode)
		return
	}
}

func (c CartController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	req, err := parseDeleteProductRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.DeleteProduct(r.Context(), req.UserId, req.Sku)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
