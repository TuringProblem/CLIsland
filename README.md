# CLIsland

CLIsland is a command-line, single-player narrative game inspired by Love Island, written in Go. The game features branching dialogue, relationship management, and drama, all through a clean, modular, and testable architecture.

## Features
- **Clean, Functional Architecture**: Modular, testable, and extensible codebase using domain-driven design principles.
- **Gameplay Loop**: Narrative-driven, state-machine-based gameplay with choices, relationships, and drama.
- **Separation of Concerns**: Clear separation between game engine, I/O, state, and scripting.
- **Testability**: Interfaces and dependency injection for easy unit testing.
- **DevOps-Friendly**: Linting, formatting, test coverage, and CI setup.
- **Configurable**: Characters, events, and dialogue trees defined via config files (JSON/YAML).
- **Command-line UI**: Minimal, testable CLI (TUI support planned).

## Current Status ✅
- **Core Architecture**: Complete with domain models, interfaces, and services
- **Game Engine**: Fully functional with stub implementations
- **CLI Interface**: Working game loop with character interactions and events
- **Testing**: Unit tests for core functionality
- **Build System**: Makefile and CI setup ready

## Getting Started

### Prerequisites
- Go 1.23+
- (Optional) [golangci-lint](https://golangci-lint.run/)

### Quick Start
```bash
# Build the game
make build

# Run the game
./clisland
```

### Running the Game
```bash
# Option 1: Build and run
go build -o clisland ./cmd/main.go
./clisland

# Option 2: Run directly
go run ./cmd/main.go
```

### Building
```bash
# Using make
make build

# Using go directly
go build -o clisland ./cmd/main.go
```

### Testing
```bash
# Run all tests
go test ./...

# Run specific test package
go test ./internal/services/

# Run with coverage
make coverage
```

### Formatting
```bash
# Format the code
make fmt
```

## How to Play

1. **Start the Game**: Run `./clisland` and enter your name
2. **Main Menu Options**:
   - **Handle Current Event**: Make choices in story events
   - **Interact with Character**: Talk, date, or challenge other contestants
   - **Advance to Next Day**: Progress the story
   - **View Detailed Stats**: Check your relationships and progress
   - **Save Game**: Save your progress
   - **Quit**: Exit the game

3. **Gameplay Elements**:
   - **Events**: Story-driven scenarios with multiple choices
   - **Characters**: 3 initial contestants (Emma, James, Sophie) with unique personalities
   - **Relationships**: Build affection and trust with characters
   - **Stats**: Manage energy, confidence, popularity, and money
   - **Days**: 30-day game cycle with different event types

## Project Structure
```
CLIsland/
├── cmd/
│   └── main.go              # CLI entrypoint and game loop
├── internal/
│   ├── domain/
│   │   ├── models.go        # Core domain models (Player, Character, Event, etc.)
│   │   └── interfaces.go    # Service interfaces for dependency injection
│   ├── services/
│   │   ├── game_engine.go   # Core game logic implementation
│   │   ├── stub_services.go # Stub implementations for testing
│   │   └── game_engine_test.go # Unit tests
│   └── repositories/
│       └── memory_repository.go # In-memory data storage
├── data/                    # Game data (names, config, etc.)
├── utils/                   # Utility functions
├── scripts/                 # DevOps and helper scripts
└── Makefile                 # Build, test, and format commands
```

## Architecture

### Domain Models
- **Player**: Main character with stats, personality, and relationships
- **Character**: Other contestants with unique traits and stats
- **Event**: Story events with choices and effects
- **Relationship**: Dynamic connections between characters
- **Effect**: Changes to game state (stats, relationships, etc.)

### Services
- **GameEngine**: Core game logic and flow control
- **EventManager**: Event creation and management
- **CharacterManager**: Character operations
- **RelationshipManager**: Relationship dynamics and interactions
- **EffectProcessor**: Applying effects to game state
- **RequirementChecker**: Validating event/choice requirements
- **ConfigProvider**: Game configuration and data

### Repositories
- **MemoryStateRepository**: In-memory game state storage
- **MemoryEventRepository**: In-memory event storage
- **MemoryCharacterRepository**: In-memory character storage

## Development Tools
- **Makefile**: Common tasks for build, test, and coverage
- **Testing**: Unit tests with good coverage of core functionality

## Extending the Game

### Adding New Characters
Edit `internal/services/stub_services.go` in the `GetCharacterConfigs` method:
```go
{
    ID:   "char_4",
    Name: "NewCharacter",
    Age:  25,
    Personality: domain.Personality{
        Openness: 60.0,
        // ... other traits
    },
    // ... other properties
}
```

### Adding New Events
Edit `internal/services/stub_services.go` in the `GetEventConfigs` method:
```go
{
    ID:          "event_3",
    Title:       "New Event",
    Description: "Description of the new event",
    Type:        domain.EventTypeDrama,
    Choices: []domain.Choice{
        // ... choices with effects
    },
    IsActive: true,
}
```

### Adding New Effects
Extend the `EffectProcessor` in `internal/services/stub_services.go`:
```go
case domain.EffectTypeNewEffect:
    // Handle new effect type
```

## Testing Strategy
- **Unit Tests**: Test individual services and components
- **Integration Tests**: Test service interactions
- **Mock Services**: Use stub implementations for isolated testing
- **Coverage**: Aim for >80% test coverage

## License
MIT
