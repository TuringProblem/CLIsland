# Module Organization Guide

This document outlines the recommended structure and organization for the CLIsland project modules.

## Project Structure

```
CLIsland/
├── cmd/                    # Command-line applications
│   ├── main.go            # Main application entry point
│   ├── person.go          # Person-related types and functions
│   ├── colors.go          # Color constants for UI
│   ├── constants.go       # General constants
│   ├── home.go            # Home screen logic
│   ├── prompts.go         # User prompts and menus
│   └── tag.go             # UI tag functions
├── internal/              # Private application code
│   ├── domain/            # Domain models and interfaces
│   │   ├── models.go      # Core domain models
│   │   └── interfaces.go  # Service interfaces
│   ├── services/          # Business logic services
│   │   ├── game_engine.go # Core game logic
│   │   └── stub_services.go # Stub implementations
│   └── repositories/      # Data access layer
│       └── memory_repository.go
├── utils/                 # Shared utility functions
│   ├── character_generator.go # Character generation
│   ├── person_generator.go    # Person generation
│   ├── filters.go         # Data filtering utilities
│   └── README.md          # Utils documentation
├── data/                  # Static data files
│   └── names/             # Name lists
│       ├── boys.txt
│       └── girls.txt
├── tests/                 # Test files (organized by type)
│   ├── unit/              # Unit tests
│   │   └── utils_test.go
│   ├── integration/       # Integration tests
│   ├── e2e/              # End-to-end tests
│   ├── test_config.go    # Test configuration
│   └── run_tests.sh      # Test runner script
├── docs/                  # Documentation
│   └── MODULE_ORGANIZATION.md
├── scripts/               # Build and deployment scripts
├── build/                 # Build artifacts
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── Makefile               # Build automation
├── README.md              # Project overview
└── LICENSE                # License file
```

## Module Organization Principles

### 1. Separation of Concerns

- **cmd/**: Contains only the application entry points and CLI-specific code
- **internal/**: Contains all private application logic that shouldn't be imported by other projects
- **utils/**: Contains shared utility functions that could potentially be used by other projects
- **data/**: Contains static data files used by the application

### 2. Package Naming Conventions

- Use lowercase, single-word package names
- Avoid underscores or mixed caps
- Use descriptive names that indicate the package's purpose

### 3. File Organization

- Group related functionality in the same package
- Keep files focused on a single responsibility
- Use consistent naming patterns within packages

## Testing Strategy

### Test Organization

```
tests/
├── unit/                  # Unit tests for individual functions
├── integration/           # Integration tests for component interaction
├── e2e/                  # End-to-end tests for full workflows
├── test_config.go        # Shared test configuration
└── run_tests.sh          # Test execution script
```

### Test Naming Conventions

- Unit tests: `TestFunctionName`
- Integration tests: `TestComponent_Scenario`
- E2E tests: `TestWorkflow_EndToEnd`

### Running Tests

```bash
# Run all tests
make test-all

# Run specific test types
make test-unit
make test-integration
make test-e2e

# Run tests with custom configuration
TEST_VERBOSE=true make test-unit
TEST_RUN_SLOW=true make test-all
```

## Module Dependencies

### Dependency Flow

```
cmd/ → internal/ → utils/
     ↓
   data/
```

### Import Rules

1. **cmd/ packages** can import:
   - internal/ packages
   - utils/ packages
   - Standard library packages

2. **internal/ packages** can import:
   - Other internal/ packages
   - utils/ packages
   - Standard library packages

3. **utils/ packages** can import:
   - Other utils/ packages
   - Standard library packages
   - External dependencies

4. **tests/ packages** can import:
   - All application packages
   - Testing packages

## Adding New Modules

### 1. Identify the Module Type

- **Business Logic**: Place in `internal/services/`
- **Data Models**: Place in `internal/domain/`
- **Data Access**: Place in `internal/repositories/`
- **Utilities**: Place in `utils/`
- **CLI Commands**: Place in `cmd/`

### 2. Create the Module Structure

```go
// Example: internal/services/new_service.go
package services

import (
    "context"
    "github.com/TuringProblem/CLIsland/internal/domain"
)

type NewService struct {
    // Dependencies
}

func NewNewService() *NewService {
    return &NewService{}
}

func (s *NewService) DoSomething(ctx context.Context) error {
    // Implementation
    return nil
}
```

### 3. Add Tests

```go
// Example: tests/unit/new_service_test.go
package tests

import (
    "testing"
    "github.com/TuringProblem/CLIsland/internal/services"
)

func TestNewService_DoSomething(t *testing.T) {
    service := services.NewNewService()
    
    err := service.DoSomething(context.Background())
    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }
}
```

### 4. Update Documentation

- Add module description to relevant README files
- Update this guide if new patterns are established

## Best Practices

### 1. Interface Design

- Define interfaces in the package that uses them
- Keep interfaces small and focused
- Use interfaces for dependency injection

### 2. Error Handling

- Return errors rather than panicking
- Use wrapped errors for context
- Define custom error types when needed

### 3. Configuration

- Use environment variables for configuration
- Provide sensible defaults
- Validate configuration on startup

### 4. Logging

- Use structured logging
- Include context in log messages
- Use appropriate log levels

### 5. Documentation

- Document all exported functions
- Include usage examples
- Keep documentation up to date

## Migration Guide

### Moving from Old Structure

If you have existing code that doesn't follow this structure:

1. **Identify the module type** based on its purpose
2. **Move the code** to the appropriate directory
3. **Update imports** in all affected files
4. **Add tests** in the appropriate test directory
5. **Update documentation** to reflect the new structure

### Example Migration

```bash
# Before
src/
├── character.go
├── character_test.go
└── main.go

# After
utils/
├── character_generator.go
└── README.md
tests/
└── unit/
    └── utils_test.go
cmd/
└── main.go
```

## Tools and Scripts

### Makefile Targets

- `make build`: Build the application
- `make test`: Run unit tests
- `make test-all`: Run all tests
- `make fmt`: Format code
- `make coverage`: Generate coverage report
- `make clean`: Clean build artifacts

### Test Scripts

- `./tests/run_tests.sh unit`: Run unit tests
- `./tests/run_tests.sh integration`: Run integration tests
- `./tests/run_tests.sh e2e`: Run end-to-end tests
- `./tests/run_tests.sh all`: Run all tests

## Conclusion

This organization structure promotes:

- **Maintainability**: Clear separation of concerns
- **Testability**: Organized test structure
- **Scalability**: Easy to add new modules
- **Reusability**: Shared utilities in utils/
- **Clarity**: Consistent naming and structure

Follow these guidelines to maintain a clean, organized codebase that's easy to understand and extend. 