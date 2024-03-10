package cart

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

type validationRequestGet struct {
	UserId int64 `validate:"gt=0"`
}

type validationRequestAddProduct struct {
	UserId int64  `validate:"gt=0"`
	Sku    int64  `validate:"gt=0"`
	Count  uint64 `validate:"gt=0"`
}

type validationRequestDeleteProduct struct {
	UserId int64 `validate:"gt=0"`
	Sku    int64 `validate:"gt=0"`
}

type validationRequestClear struct {
	UserId int64 `validate:"gt=0"`
}

func parseGetRequest(r *http.Request) (*parsedRequestGet, error) {
	userIdPathValue := r.PathValue("user_id")
	userId, err := strconv.ParseInt(userIdPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	requestValidation := validationRequestGet{
		UserId: userId,
	}

	validate := validator.New()
	err = validate.Struct(requestValidation)
	if err != nil {
		return nil, fmt.Errorf("failed validate.Struct: %w", err)
	}

	return &parsedRequestGet{
		UserId: userId,
	}, nil
}

func parseAddProductRequest(r *http.Request) (*parsedRequestAddProduct, error) {
	userIdPathValue := r.PathValue("user_id")
	skuPathValue := r.PathValue("sku")

	userId, err := strconv.ParseInt(userIdPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	sku, err := strconv.ParseInt(skuPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	var body addProductRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		return nil, fmt.Errorf("failed json.NewDecoder: %w", err)
	}

	requestValidation := validationRequestAddProduct{
		UserId: userId,
		Sku:    sku,
		Count:  body.Count,
	}

	validate := validator.New()
	err = validate.Struct(requestValidation)
	if err != nil {
		return nil, fmt.Errorf("failed validate.Struct: %w", err)
	}

	return &parsedRequestAddProduct{
		UserId: userId,
		Sku:    sku,
		Count:  body.Count,
	}, nil
}

func parseClearRequest(r *http.Request) (*parsedRequestClear, error) {
	userIdPathValue := r.PathValue("user_id")
	userId, err := strconv.ParseInt(userIdPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	requestValidation := validationRequestClear{
		UserId: userId,
	}

	validate := validator.New()
	err = validate.Struct(requestValidation)
	if err != nil {
		return nil, fmt.Errorf("failed validate.Struct: %w", err)
	}

	return &parsedRequestClear{UserId: userId}, nil
}

func parseDeleteProductRequest(r *http.Request) (*parsedRequestDeleteProduct, error) {
	userIdPathValue := r.PathValue("user_id")
	skuPathValue := r.PathValue("sku")

	userId, err := strconv.ParseInt(userIdPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	sku, err := strconv.ParseInt(skuPathValue, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed strconv.ParseInt: %w", err)
	}

	requestValidation := validationRequestDeleteProduct{
		UserId: userId,
		Sku:    sku,
	}

	validate := validator.New()
	err = validate.Struct(requestValidation)
	if err != nil {
		return nil, fmt.Errorf("failed validate.Struct: %w", err)
	}

	return &parsedRequestDeleteProduct{UserId: userId, Sku: sku}, nil
}
