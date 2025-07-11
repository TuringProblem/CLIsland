#!/bin/bash

echo "🏝️  CLIsland Demo Script 🏝️"
echo "================================"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.20+ first."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Build the game
echo ""
echo "🔨 Building CLIsland..."
if go build -o clisland ./cmd/main.go; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

# Run tests
echo ""
echo "🧪 Running tests..."
if go test ./internal/services/ -v; then
    echo "✅ Tests passed!"
else
    echo "❌ Tests failed!"
    exit 1
fi

# Check if binary exists
if [ -f "./clisland" ]; then
    echo ""
    echo "🎮 CLIsland is ready to play!"
    echo ""
    echo "To start the game, run:"
    echo "  ./clisland"
    echo ""
    echo "Or use make commands:"
    echo "  make build    # Build the game"
    echo "  make test     # Run tests"
    echo "  make lint     # Run linter"
    echo "  make coverage # Run tests with coverage"
    echo ""
    echo "Game features:"
    echo "  - 3 characters with unique personalities"
    echo "  - Story events with multiple choices"
    echo "  - Relationship building system"
    echo "  - 30-day game cycle"
    echo "  - Stats management (energy, confidence, popularity, money)"
    echo ""
    echo "Would you like to start the game now? (y/n)"
    read -r response
    if [[ "$response" =~ ^[Yy]$ ]]; then
        echo "Starting CLIsland..."
        ./clisland
    fi
else
    echo "❌ Binary not found after build!"
    exit 1
fi 