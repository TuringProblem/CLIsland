# Makefile for CLIsland

BINARY=clisland
PKG=./...
VERSION?=$(shell git describe --tags --always --dirty)

.PHONY: all build run test test-unit test-integration test-e2e test-all fmt coverage clean lint security release help

all: build

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	go build -ldflags="-s -w -X main.version=$(VERSION)" -o $(BINARY) ./cmd/main.go

build-all: ## Build for all platforms
	@echo "Building for multiple platforms..."
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-linux-amd64 ./cmd/main.go
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-linux-arm64 ./cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-darwin-amd64 ./cmd/main.go
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-darwin-arm64 ./cmd/main.go
	GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-windows-amd64.exe ./cmd/main.go
	GOOS=windows GOARCH=arm64 go build -ldflags="-s -w -X main.version=$(VERSION)" -o clisland-windows-arm64.exe ./cmd/main.go

run: ## Run the application
	go run ./cmd/main.go

# Test targets
test: test-unit ## Run unit tests

test-unit: ## Run unit tests only
	@echo "Running unit tests..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./tests/unit/...

test-integration: ## Run integration tests
	@echo "Running integration tests..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./tests/integration/...

test-e2e: ## Run end-to-end tests
	@echo "Running end-to-end tests..."
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./tests/e2e/...

test-all: test-unit test-integration test-e2e ## Run all tests

test-coverage: ## Run tests with coverage report
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./tests/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

# Legacy test target for backward compatibility
test-legacy: ## Run legacy tests
	go test -v $(PKG)

fmt: ## Format code
	gofmt -s -w .
	goimports -w .

lint: ## Run linter
	golangci-lint run

lint-fix: ## Run linter with auto-fix
	golangci-lint run --fix

security: ## Run security scan
	gosec ./...

coverage: ## Generate coverage report
	go test -coverprofile=coverage.out $(PKG)
	go tool cover -func=coverage.out

# Clean build artifacts
clean: ## Clean build artifacts
	rm -f $(BINARY)
	rm -f clisland-*
	rm -f coverage.out
	rm -f coverage.html
	rm -f checksums.txt
	go clean -cache

# Release targets
release: clean build-all ## Build release artifacts
	@echo "Creating release artifacts..."
	sha256sum clisland-* > checksums.txt
	@echo "Release artifacts created:"
	@ls -la clisland-*
	@echo "Checksums:"
	@cat checksums.txt

# Development targets
dev-setup: ## Setup development environment
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# CI/CD targets
ci: test-unit lint security ## Run CI checks locally 