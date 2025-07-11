# Makefile for CLIsland

BINARY=clisland
PKG=./...

.PHONY: all build run test lint fmt coverage

all: build

build:
	go build -o $(BINARY) ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v $(PKG)

lint:
	golangci-lint run

fmt:
	gofmt -s -w .

coverage:
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -func=coverage.out 