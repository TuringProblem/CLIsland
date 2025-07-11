package services

import (
	"context"
	"testing"

	"github.com/TuringProblem/CLIsland/internal/repositories"
)

func TestGameEngine_StartGame(t *testing.T) {
	// Setup
	stateRepo := repositories.NewMemoryStateRepository()
	eventManager := NewStubEventManager()
	characterManager := NewStubCharacterManager()
	relationshipManager := NewStubRelationshipManager()
	effectProcessor := NewStubEffectProcessor()
	requirementChecker := NewStubRequirementChecker()
	configProvider := NewStubConfigProvider()

	gameEngine := NewGameEngineService(
		stateRepo,
		eventManager,
		characterManager,
		relationshipManager,
		effectProcessor,
		requirementChecker,
		configProvider,
	)

	ctx := context.Background()

	// Test
	gameState, err := gameEngine.StartGame(ctx, "TestPlayer")

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if gameState == nil {
		t.Fatal("Expected game state, got nil")
	}

	if gameState.Player == nil {
		t.Fatal("Expected player, got nil")
	}

	if gameState.Player.Name != "TestPlayer" {
		t.Errorf("Expected player name 'TestPlayer', got '%s'", gameState.Player.Name)
	}

	if gameState.GameDay != 1 {
		t.Errorf("Expected game day 1, got %d", gameState.GameDay)
	}

	if gameState.IsGameOver {
		t.Error("Expected game not to be over")
	}

	if len(gameState.Characters) == 0 {
		t.Error("Expected characters to be loaded")
	}

	if len(gameState.Events) == 0 {
		t.Error("Expected events to be loaded")
	}
}

func TestGameEngine_AdvanceDay(t *testing.T) {
	// Setup
	stateRepo := repositories.NewMemoryStateRepository()
	eventManager := NewStubEventManager()
	characterManager := NewStubCharacterManager()
	relationshipManager := NewStubRelationshipManager()
	effectProcessor := NewStubEffectProcessor()
	requirementChecker := NewStubRequirementChecker()
	configProvider := NewStubConfigProvider()

	gameEngine := NewGameEngineService(
		stateRepo,
		eventManager,
		characterManager,
		relationshipManager,
		effectProcessor,
		requirementChecker,
		configProvider,
	)

	ctx := context.Background()

	// Start game first
	_, err := gameEngine.StartGame(ctx, "TestPlayer")
	if err != nil {
		t.Fatalf("Failed to start game: %v", err)
	}

	// Test advance day
	gameState, err := gameEngine.AdvanceDay(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if gameState.GameDay != 2 {
		t.Errorf("Expected game day 2, got %d", gameState.GameDay)
	}

	if gameState.Player.Stats.DayNumber != 2 {
		t.Errorf("Expected player day number 2, got %d", gameState.Player.Stats.DayNumber)
	}

	// Energy should be regenerated
	if gameState.Player.Stats.Energy < 100.0 {
		t.Errorf("Expected energy to be regenerated, got %.1f", gameState.Player.Stats.Energy)
	}
}

func TestGameEngine_GetAvailableCharacters(t *testing.T) {
	// Setup
	stateRepo := repositories.NewMemoryStateRepository()
	eventManager := NewStubEventManager()
	characterManager := NewStubCharacterManager()
	relationshipManager := NewStubRelationshipManager()
	effectProcessor := NewStubEffectProcessor()
	requirementChecker := NewStubRequirementChecker()
	configProvider := NewStubConfigProvider()

	gameEngine := NewGameEngineService(
		stateRepo,
		eventManager,
		characterManager,
		relationshipManager,
		effectProcessor,
		requirementChecker,
		configProvider,
	)

	ctx := context.Background()

	// Start game first
	_, err := gameEngine.StartGame(ctx, "TestPlayer")
	if err != nil {
		t.Fatalf("Failed to start game: %v", err)
	}

	// Test get available characters
	characters, err := gameEngine.GetAvailableCharacters(ctx)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(characters) == 0 {
		t.Error("Expected available characters, got none")
	}

	// Check that all characters are available
	for _, character := range characters {
		if !character.IsAvailable {
			t.Errorf("Expected character %s to be available", character.Name)
		}
	}
}
