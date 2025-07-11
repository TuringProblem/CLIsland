package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/TuringProblem/CLIsland/internal/domain"
	"github.com/TuringProblem/CLIsland/internal/repositories"
	"github.com/TuringProblem/CLIsland/internal/services"
)

func main() {
	// Initialize all services with stubbed implementations
	stateRepo := repositories.NewMemoryStateRepository()
	eventManager := services.NewStubEventManager()
	characterManager := services.NewStubCharacterManager()
	relationshipManager := services.NewStubRelationshipManager()
	effectProcessor := services.NewStubEffectProcessor()
	requirementChecker := services.NewStubRequirementChecker()
	configProvider := services.NewStubConfigProvider()

	// Create game engine
	gameEngine := services.NewGameEngineService(
		stateRepo,
		eventManager,
		characterManager,
		relationshipManager,
		effectProcessor,
		requirementChecker,
		configProvider,
	)

	// Start the game
	runGame(gameEngine)
}

func runGame(gameEngine domain.GameEngine) {
	ctx := context.Background()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🏝️  Welcome to CLIsland! 🏝️")
	fmt.Println("A Love Island-inspired command-line adventure!")
	fmt.Println()

	// Check if there's a saved game
	_, err := gameEngine.GetCurrentState(ctx)
	if err != nil {
		// Start new game
		fmt.Print("Enter your name: ")
		scanner.Scan()
		playerName := strings.TrimSpace(scanner.Text())
		if playerName == "" {
			playerName = "Player"
		}

		_, err = gameEngine.StartGame(ctx, playerName)
		if err != nil {
			fmt.Printf("Error starting game: %v\n", err)
			return
		}
		fmt.Printf("Welcome to the villa, %s! Let's find love! 💕\n", playerName)
	} else {
		fmt.Println("Loading saved game...")
	}

	// Main game loop
	for {
		gameState, err := gameEngine.GetCurrentState(ctx)
		if err != nil {
			fmt.Printf("Error getting game state: %v\n", err)
			return
		}

		if gameState.IsGameOver {
			displayGameOver(gameState)
			break
		}

		displayGameState(gameState)
		displayMainMenu()

		fmt.Print("Enter your choice: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			handleCurrentEvent(ctx, gameEngine, scanner)
		case "2":
			handleCharacterInteraction(ctx, gameEngine, scanner)
		case "3":
			handleAdvanceDay(ctx, gameEngine)
		case "4":
			handleViewStats(gameState)
		case "5":
			handleSaveGame(ctx, gameEngine)
		case "6":
			fmt.Println("Thanks for playing CLIsland! 👋")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		fmt.Println()
	}
}

func displayGameState(gameState *domain.GameState) {
	fmt.Printf("\n=== Day %d ===\n", gameState.GameDay)
	fmt.Printf("Player: %s\n", gameState.Player.Name)
	fmt.Printf("Energy: %.1f%% | Confidence: %.1f%% | Popularity: %.1f%% | Money: $%d\n",
		gameState.Player.Stats.Energy,
		gameState.Player.Stats.Confidence,
		gameState.Player.Stats.Popularity,
		gameState.Player.Stats.Money,
	)

	if gameState.CurrentEvent != nil {
		fmt.Printf("\n📢 Current Event: %s\n", gameState.CurrentEvent.Title)
		fmt.Printf("   %s\n", gameState.CurrentEvent.Description)
	}

	fmt.Printf("\n👥 Available Characters (%d):\n", len(gameState.Characters))
	for _, character := range gameState.Characters {
		if character.IsAvailable {
			relationship, exists := gameState.Player.Relationships[character.ID]
			affection := 0.0
			if exists {
				affection = relationship.Affection
			}
			fmt.Printf("   %s (Age: %d) - Affection: %.1f\n", character.Name, character.Age, affection)
		}
	}
}

func displayMainMenu() {
	fmt.Println("\n=== Main Menu ===")
	fmt.Println("1. Handle current event")
	fmt.Println("2. Interact with character")
	fmt.Println("3. Advance to next day")
	fmt.Println("4. View detailed stats")
	fmt.Println("5. Save game")
	fmt.Println("6. Quit")
}

func handleCurrentEvent(ctx context.Context, gameEngine domain.GameEngine, scanner *bufio.Scanner) {
	gameState, err := gameEngine.GetCurrentState(ctx)
	if err != nil {
		fmt.Printf("Error getting game state: %v\n", err)
		return
	}

	if gameState.CurrentEvent == nil {
		fmt.Println("No current event to handle.")
		return
	}

	fmt.Printf("\n=== %s ===\n", gameState.CurrentEvent.Title)
	fmt.Printf("%s\n\n", gameState.CurrentEvent.Description)

	fmt.Println("Available choices:")
	for i, choice := range gameState.CurrentEvent.Choices {
		fmt.Printf("%d. %s\n", i+1, choice.Text)
		fmt.Printf("   %s\n", choice.Description)
	}

	fmt.Print("\nEnter your choice (number): ")
	scanner.Scan()
	choiceStr := strings.TrimSpace(scanner.Text())
	choiceIndex, err := strconv.Atoi(choiceStr)
	if err != nil || choiceIndex < 1 || choiceIndex > len(gameState.CurrentEvent.Choices) {
		fmt.Println("Invalid choice.")
		return
	}

	selectedChoice := gameState.CurrentEvent.Choices[choiceIndex-1]
	fmt.Printf("\nYou chose: %s\n", selectedChoice.Text)

	// Process the choice
	_, err = gameEngine.ProcessChoice(ctx, selectedChoice.ID)
	if err != nil {
		fmt.Printf("Error processing choice: %v\n", err)
		return
	}

	fmt.Println("Choice processed successfully!")
}

func handleCharacterInteraction(ctx context.Context, gameEngine domain.GameEngine, scanner *bufio.Scanner) {
	characters, err := gameEngine.GetAvailableCharacters(ctx)
	if err != nil {
		fmt.Printf("Error getting characters: %v\n", err)
		return
	}

	if len(characters) == 0 {
		fmt.Println("No characters available for interaction.")
		return
	}

	fmt.Println("\n=== Character Interaction ===")
	fmt.Println("Available characters:")
	for i, character := range characters {
		fmt.Printf("%d. %s (Age: %d)\n", i+1, character.Name, character.Age)
	}

	fmt.Print("Choose a character (number): ")
	scanner.Scan()
	charChoice := strings.TrimSpace(scanner.Text())
	charIndex, err := strconv.Atoi(charChoice)
	if err != nil || charIndex < 1 || charIndex > len(characters) {
		fmt.Println("Invalid character choice.")
		return
	}

	selectedCharacter := characters[charIndex-1]

	fmt.Println("\nInteraction types:")
	fmt.Println("1. Conversation")
	fmt.Println("2. Date")
	fmt.Println("3. Challenge")
	fmt.Println("4. Gift")
	fmt.Println("5. Argument")

	fmt.Print("Choose interaction type (number): ")
	scanner.Scan()
	interactionChoice := strings.TrimSpace(scanner.Text())
	interactionIndex, err := strconv.Atoi(interactionChoice)
	if err != nil || interactionIndex < 1 || interactionIndex > 5 {
		fmt.Println("Invalid interaction choice.")
		return
	}

	var interactionType domain.InteractionType
	switch interactionIndex {
	case 1:
		interactionType = domain.InteractionTypeConversation
	case 2:
		interactionType = domain.InteractionTypeDate
	case 3:
		interactionType = domain.InteractionTypeChallenge
	case 4:
		interactionType = domain.InteractionTypeGift
	case 5:
		interactionType = domain.InteractionTypeArgument
	}

	// Perform interaction
	_, err = gameEngine.InteractWithCharacter(ctx, selectedCharacter.ID, interactionType)
	if err != nil {
		fmt.Printf("Error during interaction: %v\n", err)
		return
	}

	fmt.Printf("You had a %s with %s!\n", strings.ToLower(string(interactionType)), selectedCharacter.Name)
}

func handleAdvanceDay(ctx context.Context, gameEngine domain.GameEngine) {
	fmt.Println("Advancing to the next day...")

	gameState, err := gameEngine.AdvanceDay(ctx)
	if err != nil {
		fmt.Printf("Error advancing day: %v\n", err)
		return
	}

	if gameState.IsGameOver {
		displayGameOver(gameState)
		return
	}

	fmt.Printf("Welcome to Day %d! 🌅\n", gameState.GameDay)
	if gameState.CurrentEvent != nil {
		fmt.Printf("New event: %s\n", gameState.CurrentEvent.Title)
	}
}

func handleViewStats(gameState *domain.GameState) {
	fmt.Println("\n=== Detailed Stats ===")
	fmt.Printf("Player: %s\n", gameState.Player.Name)
	fmt.Printf("Age: %d\n", gameState.Player.Age)
	fmt.Printf("Day: %d\n", gameState.GameDay)
	fmt.Printf("Location: %s\n", gameState.Player.CurrentLocation)

	fmt.Println("\nStats:")
	fmt.Printf("  Energy: %.1f%%\n", gameState.Player.Stats.Energy)
	fmt.Printf("  Confidence: %.1f%%\n", gameState.Player.Stats.Confidence)
	fmt.Printf("  Popularity: %.1f%%\n", gameState.Player.Stats.Popularity)
	fmt.Printf("  Money: $%d\n", gameState.Player.Stats.Money)

	fmt.Println("\nPersonality:")
	fmt.Printf("  Openness: %.1f%%\n", gameState.Player.Personality.Openness)
	fmt.Printf("  Conscientiousness: %.1f%%\n", gameState.Player.Personality.Conscientiousness)
	fmt.Printf("  Extraversion: %.1f%%\n", gameState.Player.Personality.Extraversion)
	fmt.Printf("  Agreeableness: %.1f%%\n", gameState.Player.Personality.Agreeableness)
	fmt.Printf("  Neuroticism: %.1f%%\n", gameState.Player.Personality.Neuroticism)

	fmt.Println("\nRelationships:")
	for characterID, relationship := range gameState.Player.Relationships {
		character, exists := gameState.Characters[characterID]
		if exists {
			fmt.Printf("  %s: Affection %.1f, Trust %.1f, Status: %s\n",
				character.Name,
				relationship.Affection,
				relationship.Trust,
				relationship.Status,
			)
		}
	}

	fmt.Printf("\nInventory (%d items):\n", len(gameState.Player.Inventory))
	for _, item := range gameState.Player.Inventory {
		fmt.Printf("  %s: %s (Value: $%d)\n", item.Name, item.Description, item.Value)
	}
}

func handleSaveGame(ctx context.Context, gameEngine domain.GameEngine) {
	err := gameEngine.SaveGame(ctx)
	if err != nil {
		fmt.Printf("Error saving game: %v\n", err)
		return
	}
	fmt.Println("Game saved successfully! 💾")
}

func displayGameOver(gameState *domain.GameState) {
	fmt.Println("\n=== GAME OVER ===")
	fmt.Printf("Congratulations! You've completed CLIsland!\n")
	fmt.Printf("Final day: %d\n", gameState.GameDay)

	if gameState.Winner != "" {
		if character, exists := gameState.Characters[gameState.Winner]; exists {
			fmt.Printf("Winner: %s! 💕\n", character.Name)
		}
	}

	fmt.Println("\nFinal Stats:")
	handleViewStats(gameState)
}
