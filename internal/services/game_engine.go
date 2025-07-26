package services

import (
	"context"
	"fmt"
	"time"

	"github.com/TuringProblem/CLIsland/internal/domain"
)

type GameEngineService struct {
	stateRepo           domain.StateRepository
	eventManager        domain.EventManager
	characterManager    domain.CharacterManager
	relationshipManager domain.RelationshipManager
	effectProcessor     domain.EffectProcessor
	requirementChecker  domain.RequirementChecker
	configProvider      domain.ConfigProvider
}

func NewGameEngineService(
	stateRepo domain.StateRepository,
	eventManager domain.EventManager,
	characterManager domain.CharacterManager,
	relationshipManager domain.RelationshipManager,
	effectProcessor domain.EffectProcessor,
	requirementChecker domain.RequirementChecker,
	configProvider domain.ConfigProvider,
) *GameEngineService {
	return &GameEngineService{
		stateRepo:           stateRepo,
		eventManager:        eventManager,
		characterManager:    characterManager,
		relationshipManager: relationshipManager,
		effectProcessor:     effectProcessor,
		requirementChecker:  requirementChecker,
		configProvider:      configProvider,
	}
}

func (g *GameEngineService) StartGame(ctx context.Context, playerName string) (*domain.GameState, error) {
	config, err := g.configProvider.GetGameConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get game config: %w", err)
	}

	player := &domain.Player{
		ID:   generateID(),
		Name: playerName,
		Age:  25,
		Personality: domain.Personality{
			Openness:          50.0,
			Conscientiousness: 50.0,
			Extraversion:      50.0,
			Agreeableness:     50.0,
			Neuroticism:       50.0,
		},
		Relationships: make(map[string]domain.Relationship),
		Stats: domain.PlayerStats{
			Popularity: config.StartingPopularity,
			Confidence: config.StartingConfidence,
			Energy:     config.StartingEnergy,
			Money:      config.StartingMoney,
			DayNumber:  1,
		},
		Inventory:       []domain.Item{},
		CurrentLocation: "villa",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	characters, err := g.configProvider.GetCharacterConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load characters: %w", err)
	}

	characterMap := make(map[string]*domain.Character)
	for i := range characters {
		if i >= config.MaxCharacters {
			break
		}
		characters[i].ID = generateID()
		characters[i].IsAvailable = true
		characters[i].CreatedAt = time.Now()
		characterMap[characters[i].ID] = characters[i]
	}

	events, err := g.configProvider.GetEventConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load events: %w", err)
	}

	eventMap := make(map[string]*domain.Event)
	for i := range events {
		events[i].ID = generateID()
		events[i].IsActive = true
		events[i].CreatedAt = time.Now()
		eventMap[events[i].ID] = events[i]
	}

	gameState := &domain.GameState{
		Player:       player,
		Characters:   characterMap,
		Events:       eventMap,
		CurrentEvent: nil,
		GameDay:      1,
		IsGameOver:   false,
		Winner:       "",
	}

	if err := g.stateRepo.Save(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to save initial game state: %w", err)
	}

	return gameState, nil
}

func (g *GameEngineService) ProcessChoice(ctx context.Context, choiceID string) (*domain.GameState, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	if gameState.CurrentEvent == nil {
		return nil, fmt.Errorf("no current event to process choice for")
	}

	var selectedChoice *domain.Choice
	for _, choice := range gameState.CurrentEvent.Choices {
		if choice.ID == choiceID {
			selectedChoice = &choice
			break
		}
	}

	if selectedChoice == nil {
		return nil, fmt.Errorf("choice with ID %s not found", choiceID)
	}

	if len(selectedChoice.Requirements) > 0 {
		requirements := make([]*domain.Requirement, len(selectedChoice.Requirements))
		for i := range selectedChoice.Requirements {
			requirements[i] = &selectedChoice.Requirements[i]
		}
		meetsRequirements, err := g.requirementChecker.CheckRequirements(ctx, requirements, gameState)
		if err != nil {
			return nil, fmt.Errorf("failed to check requirements: %w", err)
		}
		if !meetsRequirements {
			return nil, fmt.Errorf("requirements not met for choice %s", choiceID)
		}
	}

	effects := make([]*domain.Effect, len(selectedChoice.Effects))
	for i := range selectedChoice.Effects {
		effects[i] = &selectedChoice.Effects[i]
	}
	if err := g.effectProcessor.ApplyEffects(ctx, effects, gameState); err != nil {
		return nil, fmt.Errorf("failed to apply choice effects: %w", err)
	}

	gameState.Player.UpdatedAt = time.Now()

	if err := g.stateRepo.Save(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to save game state: %w", err)
	}

	return gameState, nil
}

func (g *GameEngineService) AdvanceDay(ctx context.Context) (*domain.GameState, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	config, err := g.configProvider.GetGameConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get game config: %w", err)
	}

	if gameState.GameDay >= config.MaxDays {
		gameState.IsGameOver = true
		gameState.Winner = determineWinner(gameState)
		return gameState, nil
	}

	gameState.GameDay++
	gameState.Player.Stats.DayNumber = gameState.GameDay

	gameState.Player.Stats.Energy = min(100.0, gameState.Player.Stats.Energy+20.0)

	for _, character := range gameState.Characters {
		if character.IsAvailable {
			character.Stats.Energy = min(100.0, character.Stats.Energy+15.0)
			character.Stats.Stress = max(0.0, character.Stats.Stress-5.0)
		}
	}

	newEvent, err := g.generateDailyEvent(ctx, gameState)
	if err != nil {
		return nil, fmt.Errorf("failed to generate daily event: %w", err)
	}
	gameState.CurrentEvent = newEvent

	if err := g.stateRepo.Save(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to save game state: %w", err)
	}

	return gameState, nil
}

func (g *GameEngineService) EndGame(ctx context.Context) error {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load game state: %w", err)
	}

	gameState.IsGameOver = true
	gameState.Winner = determineWinner(gameState)

	return g.stateRepo.Save(ctx, gameState)
}

func (g *GameEngineService) GetCurrentState(ctx context.Context) (*domain.GameState, error) {
	return g.stateRepo.Load(ctx)
}

func (g *GameEngineService) SaveGame(ctx context.Context) error {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load game state: %w", err)
	}

	return g.stateRepo.Save(ctx, gameState)
}

func (g *GameEngineService) LoadGame(ctx context.Context) (*domain.GameState, error) {
	return g.stateRepo.Load(ctx)
}

func (g *GameEngineService) GetAvailableEvents(ctx context.Context) ([]*domain.Event, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	var availableEvents []*domain.Event
	for _, event := range gameState.Events {
		if event.IsActive {
			if len(event.Requirements) > 0 {
				requirements := make([]*domain.Requirement, len(event.Requirements))
				for i := range event.Requirements {
					requirements[i] = &event.Requirements[i]
				}
				meetsRequirements, err := g.requirementChecker.CheckRequirements(ctx, requirements, gameState)
				if err != nil {
					continue
				}
				if meetsRequirements {
					availableEvents = append(availableEvents, event)
				}
			} else {
				availableEvents = append(availableEvents, event)
			}
		}
	}

	return availableEvents, nil
}

func (g *GameEngineService) TriggerEvent(ctx context.Context, eventID string) (*domain.GameState, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	event, exists := gameState.Events[eventID]
	if !exists {
		return nil, fmt.Errorf("event with ID %s not found", eventID)
	}

	if !event.IsActive {
		return nil, fmt.Errorf("event %s is not active", eventID)
	}

	if len(event.Requirements) > 0 {
		requirements := make([]*domain.Requirement, len(event.Requirements))
		for i := range event.Requirements {
			requirements[i] = &event.Requirements[i]
		}
		meetsRequirements, err := g.requirementChecker.CheckRequirements(ctx, requirements, gameState)
		if err != nil {
			return nil, fmt.Errorf("failed to check event requirements: %w", err)
		}
		if !meetsRequirements {
			return nil, fmt.Errorf("requirements not met for event %s", eventID)
		}
	}

	gameState.CurrentEvent = event

	if err := g.stateRepo.Save(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to save game state: %w", err)
	}

	return gameState, nil
}

func (g *GameEngineService) GetAvailableCharacters(ctx context.Context) ([]*domain.Character, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	var availableCharacters []*domain.Character
	for _, character := range gameState.Characters {
		if character.IsAvailable {
			availableCharacters = append(availableCharacters, character)
		}
	}

	return availableCharacters, nil
}

func (g *GameEngineService) InteractWithCharacter(ctx context.Context, characterID string, interactionType domain.InteractionType) (*domain.GameState, error) {
	gameState, err := g.stateRepo.Load(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game state: %w", err)
	}

	character, exists := gameState.Characters[characterID]
	if !exists {
		return nil, fmt.Errorf("character with ID %s not found", characterID)
	}

	if !character.IsAvailable {
		return nil, fmt.Errorf("character %s is not available", characterID)
	}

	interaction := &domain.Interaction{
		Type:        interactionType,
		Description: generateInteractionDescription(interactionType, character.Name),
		Effects:     generateInteractionEffects(interactionType),
		Timestamp:   time.Now(),
	}

	if err := g.relationshipManager.AddInteraction(ctx, gameState.Player.ID, characterID, interaction); err != nil {
		return nil, fmt.Errorf("failed to add interaction: %w", err)
	}

	effects := make([]*domain.Effect, len(interaction.Effects))
	for i := range interaction.Effects {
		effects[i] = &interaction.Effects[i]
	}
	if err := g.effectProcessor.ApplyEffects(ctx, effects, gameState); err != nil {
		return nil, fmt.Errorf("failed to apply interaction effects: %w", err)
	}

	if err := g.stateRepo.Save(ctx, gameState); err != nil {
		return nil, fmt.Errorf("failed to save game state: %w", err)
	}

	return gameState, nil
}

func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func determineWinner(gameState *domain.GameState) string {
	var winner string
	maxPopularity := -1.0

	for characterID, relationship := range gameState.Player.Relationships {
		if relationship.Affection > maxPopularity {
			maxPopularity = relationship.Affection
			winner = characterID
		}
	}

	return winner
}

func (g *GameEngineService) generateDailyEvent(ctx context.Context, gameState *domain.GameState) (*domain.Event, error) {
	config, err := g.configProvider.GetGameConfig(ctx)
	if err != nil {
		return nil, err
	}

	var eventType domain.EventType
	switch {
	case gameState.GameDay == config.EliminationDay:
		eventType = domain.EventTypeElimination
	case gameState.GameDay == config.FinaleDay:
		eventType = domain.EventTypeRecoupling
	case gameState.GameDay%3 == 0:
		eventType = domain.EventTypeChallenge
	case gameState.GameDay%2 == 0:
		eventType = domain.EventTypeDate
	default:
		eventType = domain.EventTypeDrama
	}

	for _, event := range gameState.Events {
		if event.Type == eventType && event.IsActive {
			return event, nil
		}
	}

	for _, event := range gameState.Events {
		if event.IsActive {
			return event, nil
		}
	}

	return nil, fmt.Errorf("no available events found")
}

func generateInteractionDescription(interactionType domain.InteractionType, characterName string) string {
	switch interactionType {
	case domain.InteractionTypeConversation:
		return fmt.Sprintf("You had a deep conversation with %s", characterName)
	case domain.InteractionTypeDate:
		return fmt.Sprintf("You went on a romantic date with %s", characterName)
	case domain.InteractionTypeChallenge:
		return fmt.Sprintf("You participated in a challenge with %s", characterName)
	case domain.InteractionTypeGift:
		return fmt.Sprintf("You gave a thoughtful gift to %s", characterName)
	case domain.InteractionTypeArgument:
		return fmt.Sprintf("You had a heated argument with %s", characterName)
	default:
		return fmt.Sprintf("You interacted with %s", characterName)
	}
}

func generateInteractionEffects(interactionType domain.InteractionType) []domain.Effect {
	var effects []domain.Effect

	switch interactionType {
	case domain.InteractionTypeConversation:
		effects = append(effects, domain.Effect{
			Type:        domain.EffectTypeAffection,
			Target:      "player",
			Value:       5.0,
			Description: "Deep conversation increased affection",
		})
	case domain.InteractionTypeDate:
		effects = append(effects, domain.Effect{
			Type:        domain.EffectTypeAffection,
			Target:      "player",
			Value:       15.0,
			Description: "Romantic date significantly increased affection",
		})
	case domain.InteractionTypeChallenge:
		effects = append(effects, domain.Effect{
			Type:        domain.EffectTypeTrust,
			Target:      "player",
			Value:       10.0,
			Description: "Team challenge built trust",
		})
	case domain.InteractionTypeGift:
		effects = append(effects, domain.Effect{
			Type:        domain.EffectTypeAffection,
			Target:      "player",
			Value:       8.0,
			Description: "Thoughtful gift increased affection",
		})
	case domain.InteractionTypeArgument:
		effects = append(effects, domain.Effect{
			Type:        domain.EffectTypeAffection,
			Target:      "player",
			Value:       -10.0,
			Description: "Argument decreased affection",
		})
	}

	return effects
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
