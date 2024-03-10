package main

import (
	"context"
	"log"

	"route256.ozon.ru/project/cart/internal/pkg/app"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.New(ctx)
	if err != nil {
		log.Fatalf("cannot build application: %s", err)
	}

	if err = a.Run(); err != nil {
		log.Fatalf("cannot run app: %s", err)
	}
}
