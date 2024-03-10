package app

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"route256.ozon.ru/project/cart/internal/config"
)

type App struct {
	ctx     context.Context
	config  config.Config
	handler http.Handler
}

func New(ctx context.Context) (*App, error) {
	config := config.Config{
		Port:                "8082",
		Host:                "0.0.0.0",
		ProductServiceUrl:   "http://route256.pavl.uk:8080",
		ProductServiceToken: "testtoken",
	}

	mux := http.NewServeMux()
	handler := setupRouter(mux, &config)

	return &App{
		ctx:     ctx,
		config:  config,
		handler: handler,
	}, nil
}

func (a App) Run() error {
	url := net.JoinHostPort(a.config.Host, a.config.Port)
	if err := http.ListenAndServe(url, a.handler); err != nil {
		return fmt.Errorf("failed http.ListenAndServe: %w", err)
	}

	return nil
}
