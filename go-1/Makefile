SHELL := /bin/bash
GOFILES := $(shell find . -name '*.go' -type f)

build:
	go build -o kafka-api-consumer ./cmd/consumer

test:
	go test ./... -coverprofile=coverage.out

test-integration:
	go test -tags=integration ./cmd/consumer

cover:
	go tool cover -html=coverage.out

lint:
	golangci-lint run

run:
	go run ./cmd/consumer

clean:
	rm -f kafka-api-consumer coverage.out

compose-up:
	docker-compose up --build

compose-down:
	docker-compose down
