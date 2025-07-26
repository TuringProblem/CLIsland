# Test Organization Summary

This document summarizes the test organization structure implemented for the CLIsland project.

## Test Structure Overview

```
tests/
├── unit/                  # Unit tests for individual functions
│   └── utils_test.go     # Tests for utils package
├── integration/           # Integration tests (to be added)
├── e2e/                  # End-to-end tests (to be added)
├── test_config.go        # Shared test configuration
└── run_tests.sh          # Test runner script
```

## Key Benefits

### 1. **Organized Test Structure**
- **Unit Tests**: Test individual functions and methods in isolation
- **Integration Tests**: Test how components work together
- **E2E Tests**: Test complete user workflows

### 2. **Easy Test Execution**
- Use Makefile targets: `make test-unit`, `make test-integration`, `make test-e2e`
- Use test runner script: `./tests/run_tests.sh [unit|integration|e2e|all]`
- Environment variable configuration for test behavior

### 3. **Scalable Architecture**
- Easy to add new test types
- Centralized test configuration
- Consistent test patterns across the project

## Current Implementation

### Generated Functions

1. **`GenerateCharacters(characters []string, playerSex string) []string`**
   - Generates character names based on player sex
   - Male player: 5 boys + 6 girls
   - Female player: 6 boys + 5 girls
   - Returns shuffled list of names

2. **`GeneratePerson(sex string) Person`**
   - Creates complete Person objects with realistic attributes
   - Includes name, age, height, weight, sex, and interests
   - Age range: 18-35 years
   - Height: Realistic ranges for each sex
   - Weight: BMI-based calculation
   - Interests: 3-6 random interests with weights (1-10)

3. **`GeneratePersonList(playerSex string) []Person`**
   - Generates full list of Person objects
   - Same logic as GenerateCharacters but returns complete objects
   - Shuffled order for randomization

### Person Object Features

```go
type Person struct {
    Name      string
    Age       int
    Height    float64 // inches
    Weight    float64 // kg
    Sex       string
    Interests Interests
}
```

**Available Methods:**
- `GetInterests() []Interest` - Returns list of interests
- `GetInterestWeight(interest Interest) int` - Returns weight of specific interest
- `HasInterest(interest Interest) bool` - Checks if person has interest
- `GetHeightInFeet() string` - Returns height in feet/inches format
- `GetHeightInCm() float64` - Returns height in centimeters
- `GetWeightInKg() float64` - Returns weight in kilograms
- `GetWeightInLbs() float64` - Returns weight in pounds

## Usage Examples

### Basic Character Generation
```go
import "github.com/TuringProblem/CLIsland/utils"

// Generate character names
characters := utils.GenerateCharacters([]string{}, "male")
fmt.Printf("Generated %d characters: %v\n", len(characters), characters)
```

### Full Person Generation
```go
// Generate complete person objects
people := utils.GeneratePersonList("female")
for i, person := range people {
    fmt.Printf("%d. %s (%s, %d years old, %s)\n", 
        i+1, person.Name, person.Sex, person.Age, person.GetHeightInFeet())
    fmt.Printf("   Weight: %.1f kg (%.1f lbs)\n", 
        person.GetWeightInKg(), person.GetWeightInLbs())
}
```

### Individual Person Creation
```go
// Create a single person
person := utils.GeneratePerson("male")
fmt.Printf("Created: %s, %d years old, %s\n", 
    person.Name, person.Age, person.GetHeightInFeet())
```

## Running Tests

### Using Makefile
```bash
# Run unit tests only
make test-unit

# Run all tests
make test-all

# Run legacy tests (old structure)
make test-legacy
```

### Using Test Runner Script
```bash
# Run all tests
./tests/run_tests.sh all

# Run specific test types
./tests/run_tests.sh unit
./tests/run_tests.sh integration
./tests/run_tests.sh e2e
```

### Using Go Directly
```bash
# Run unit tests
go test -v ./tests/unit/...

# Run with coverage
go test -coverprofile=coverage.out ./tests/unit/...
go tool cover -func=coverage.out
```

## Test Configuration

### Environment Variables
- `TEST_VERBOSE=true` - Enable verbose test output
- `TEST_RUN_SLOW=true` - Run slow tests (skipped by default)
- `TEST_DATA_PATH=path` - Specify test data directory

### Test Utilities
```go
import "github.com/TuringProblem/CLIsland/tests"

// Skip slow tests
tests.SkipIfSlowTest(t)

// Skip if test data not available
tests.SkipIfNoTestData(t)

// Get test configuration
config := tests.GetTestConfig()
```

## Adding New Tests

### 1. Unit Tests
Create tests in `tests/unit/` for individual functions:
```go
// tests/unit/new_feature_test.go
package tests

import (
    "testing"
    "github.com/TuringProblem/CLIsland/utils"
)

func TestNewFeature(t *testing.T) {
    result := utils.NewFeature()
    if result == nil {
        t.Error("Expected result, got nil")
    }
}
```

### 2. Integration Tests
Create tests in `tests/integration/` for component interaction:
```go
// tests/integration/service_integration_test.go
package tests

import (
    "testing"
    "github.com/TuringProblem/CLIsland/internal/services"
)

func TestServiceIntegration(t *testing.T) {
    service := services.NewService()
    // Test service interactions
}
```

### 3. E2E Tests
Create tests in `tests/e2e/` for complete workflows:
```go
// tests/e2e/game_workflow_test.go
package tests

import (
    "testing"
    "github.com/TuringProblem/CLIsland/cmd"
)

func TestCompleteGameWorkflow(t *testing.T) {
    // Test complete game flow from start to finish
}
```

## Best Practices

### 1. Test Naming
- Unit tests: `TestFunctionName`
- Integration tests: `TestComponent_Scenario`
- E2E tests: `TestWorkflow_EndToEnd`

### 2. Test Organization
- Group related tests in the same file
- Use descriptive test names
- Include setup and teardown when needed

### 3. Test Data
- Use realistic test data
- Avoid hardcoded values when possible
- Use test configuration for data paths

### 4. Test Coverage
- Aim for high test coverage
- Focus on critical paths
- Test edge cases and error conditions

## Migration from Old Structure

If you have existing tests in the old structure:

1. **Move test files** to appropriate directories in `tests/`
2. **Update package declarations** to use `package tests`
3. **Update imports** to use the new structure
4. **Run tests** to ensure everything works

### Example Migration
```bash
# Before
utils/character_generator_test.go

# After
tests/unit/utils_test.go
```

## Conclusion

This test organization provides:

- **Clear separation** of test types
- **Easy execution** with multiple options
- **Scalable structure** for future growth
- **Consistent patterns** across the project
- **Better maintainability** with organized code

The structure supports the project's growth while maintaining clean, organized, and easily executable tests. 