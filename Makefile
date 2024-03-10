CART_BINARY_NAME=cart-server

build-all:
	cd cart && GOOS=linux GOARCH=amd64 go build -o bin/${CART_BINARY_NAME} ./cmd/${CART_BINARY_NAME}/main.go

run-all: 
	docker-compose up --build

run-cart:
	cd cart && go run ./cmd/${CART_BINARY_NAME}