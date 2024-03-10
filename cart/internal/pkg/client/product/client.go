package product

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Client http.Client

	token      string
	serviceUrl string
}

const (
	getProductPath = "get_product"
)

func New(serviceUrl string, token string) *Client {
	return &Client{
		Client: *http.DefaultClient,

		token:      token,
		serviceUrl: serviceUrl,
	}
}

func (c Client) GetProductBySku(ctx context.Context, sku int64) (*GetProductResponse, error) {
	url, err := url.JoinPath(c.serviceUrl, getProductPath)
	if err != nil {
		return nil, fmt.Errorf("failed url.JoinPath: %w", err)
	}

	req := GetProductRequest{
		Token: c.token,
		Sku:   sku,
	}

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(req)
	if err != nil {
		return nil, fmt.Errorf("failed json.NewEncoder: %w", err)
	}

	resp, err := c.Client.Post(url, "applications/json", &buf)
	if err != nil {
		return nil, fmt.Errorf("failed c.Client.Post: %w", err)
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll: %w", err)
	}

	var response GetProductResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal: %w", err)
	}

	return &response, nil
}
