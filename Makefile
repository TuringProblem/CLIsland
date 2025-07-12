# CLIsland Makefile
# Build, test, and development tasks for the Love Island CLI game

# Variables
BINARY_NAME=clisland
BUILD_DIR=build
MAIN_PACKAGE=./cmd
MODULE_NAME=github.com/TuringProblem/CLIsland

# Go build flags
LDFLAGS=-ldflags "-s -w"
BUILD_FLAGS=-trimpath

# Default target
.PHONY: all
all: build

# Build the application
.PHONY: build
build: deps
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "✅ Build complete: ./$(BINARY_NAME)"

# Build for development (with debug info)
.PHONY: build-dev
build-dev: deps
	@echo "Building $(BINARY_NAME) for development..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "✅ Development build complete: ./$(BINARY_NAME)"

# Run the application
.PHONY: run
run: deps
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_PACKAGE)

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatting complete"

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	go vet ./...
	@echo "✅ Code vetting complete"

# Run all quality checks
.PHONY: check
check: fmt vet test
	@echo "✅ All quality checks passed"

# Install the binary
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME)..."
	go install $(MAIN_PACKAGE)
	@echo "✅ $(BINARY_NAME) installed to GOPATH"

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy
	go mod verify
	@echo "✅ Dependencies installed and verified"

# Build for multiple platforms
.PHONY: build-all
build-all: clean
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux
	GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)
	
	# macOS
	GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)
	GOOS=darwin GOARCH=arm64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)
	
	# Windows
	GOOS=windows GOARCH=amd64 go build $(BUILD_FLAGS) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)
	
	@echo "✅ Multi-platform builds complete in $(BUILD_DIR)/"

# Show help
.PHONY: help
help:
	@echo "CLIsland Makefile - Available targets:"
	@echo ""
	@echo "  build        - Build the clisland executable (default)"
	@echo "  build-dev    - Build with debug info"
	@echo "  build-all    - Build for multiple platforms (Linux, macOS, Windows)"
	@echo "  run          - Run the application directly"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  vet          - Vet code for common issues"
	@echo "  deps         - Install and tidy dependencies"
	@echo "  clean        - Remove build artifacts"
	@echo "  check        - Run fmt, vet, and test"
	@echo "  install      - Install the binary to GOPATH"
	@echo "  help         - Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build   # Build the executable"
	@echo "  make run     # Run the game"
	@echo "  make test    # Run tests"
	@echo "  make check   # Run all quality checks"
