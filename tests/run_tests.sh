#!/bin/bash

# Test runner script for CLIsland
# Usage: ./tests/run_tests.sh [unit|integration|e2e|all]

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to run tests
run_tests() {
    local test_type=$1
    local test_path=$2
    
    print_status "Running $test_type tests..."
    
    if [ -d "$test_path" ]; then
        cd "$test_path"
        if go test -v ./...; then
            print_success "$test_type tests passed"
        else
            print_error "$test_type tests failed"
            return 1
        fi
        cd - > /dev/null
    else
        print_warning "No $test_type tests found at $test_path"
    fi
}

# Function to run all tests
run_all_tests() {
    print_status "Running all tests..."
    
    # Run unit tests
    run_tests "unit" "tests/unit"
    
    # Run integration tests
    run_tests "integration" "tests/integration"
    
    # Run e2e tests
    run_tests "e2e" "tests/e2e"
    
    print_success "All tests completed"
}

# Main script logic
case "${1:-all}" in
    "unit")
        run_tests "unit" "tests/unit"
        ;;
    "integration")
        run_tests "integration" "tests/integration"
        ;;
    "e2e")
        run_tests "e2e" "tests/e2e"
        ;;
    "all")
        run_all_tests
        ;;
    *)
        print_error "Unknown test type: $1"
        echo "Usage: $0 [unit|integration|e2e|all]"
        exit 1
        ;;
esac 