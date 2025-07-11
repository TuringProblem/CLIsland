# Makefile for CLIsland

BINARY=clisland
PKG=./...

.PHONY: all build run test fmt coverage

all: build

build:
	go build -o $(BINARY) ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test -v $(PKG)

fmt:
	gofmt -s -w .

coverage:
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -func=coverage.out 