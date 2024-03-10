
FROM golang:1.22-alpine as builder

WORKDIR /app

ARG CART_BINARY_NAME=cart-server

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o bin/${CART_BINARY_NAME} ./cmd/${CART_BINARY_NAME}/main.go

FROM alpine

COPY --from=builder /app/bin/${CART_BINARY_NAME} /${CART_BINARY_NAME}

CMD [ "./cart-server" ]
