package services

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/TuringProblem/CLIsland/internal/domain"
)

type StubEventManager struct{}

func NewStubEventManager() *StubEventManager {
	return &StubEventManager{}
}

func (s *StubEventManager) CreateEvent(ctx context.Context, event *domain.Event) error {
	return nil
}

func (s *StubEventManager) GetEvent(ctx context.Context, eventID string) (*domain.Event, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *StubEventManager) UpdateEvent(ctx context.Context, event *domain.Event) error {
	return nil
}

func (s *StubEventManager) DeleteEvent(ctx context.Context, eventID string) error {
	return nil
}

func (s *StubEventManager) ListEvents(ctx context.Context, eventType domain.EventType) ([]*domain.Event, error) {
	return []*domain.Event{}, nil
}

func (s *StubEventManager) ValidateEvent(ctx context.Context, event *domain.Event) error {
	return nil
}

type StubCharacterManager struct{}

func NewStubCharacterManager() *StubCharacterManager {
	return &StubCharacterManager{}
}

func (s *StubCharacterManager) CreateCharacter(ctx context.Context, character *domain.Character) error {
	return nil
}

func (s *StubCharacterManager) GetCharacter(ctx context.Context, characterID string) (*domain.Character, error) {
	return nil, fmt.Errorf("not implemented")
}

func (s *StubCharacterManager) UpdateCharacter(ctx context.Context, character *domain.Character) error {
	return nil
}

func (s *StubCharacterManager) DeleteCharacter(ctx context.Context, characterID string) error {
	return nil
}

func (s *StubCharacterManager) ListCharacters(ctx context.Context) ([]*domain.Character, error) {
	return []*domain.Character{}, nil
}

func (s *StubCharacterManager) UpdateCharacterStats(ctx context.Context, characterID string, stats domain.CharacterStats) error {
	return nil
}

type StubRelationshipManager struct {
	relationships map[string]map[string]*domain.Relationship
}

func NewStubRelationshipManager() *StubRelationshipManager {
	return &StubRelationshipManager{
		relationships: make(map[string]map[string]*domain.Relationship),
	}
}

func (s *StubRelationshipManager) GetRelationship(ctx context.Context, playerID, characterID string) (*domain.Relationship, error) {
	if playerRelationships, exists := s.relationships[playerID]; exists {
		if relationship, exists := playerRelationships[characterID]; exists {
			return relationship, nil
		}
	}

	relationship := &domain.Relationship{
		CharacterID:   characterID,
		Affection:     0.0,
		Trust:         0.0,
		Compatibility: 0.0,
		Status:        domain.RelationshipStatusSingle,
		History:       []domain.Interaction{},
		UpdatedAt:     time.Now(),
	}

	if s.relationships[playerID] == nil {
		s.relationships[playerID] = make(map[string]*domain.Relationship)
	}
	s.relationships[playerID][characterID] = relationship

	return relationship, nil
}

func (s *StubRelationshipManager) UpdateRelationship(ctx context.Context, playerID, characterID string, relationship *domain.Relationship) error {
	if s.relationships[playerID] == nil {
		s.relationships[playerID] = make(map[string]*domain.Relationship)
	}
	s.relationships[playerID][characterID] = relationship
	return nil
}

func (s *StubRelationshipManager) AddInteraction(ctx context.Context, playerID, characterID string, interaction *domain.Interaction) error {
	relationship, err := s.GetRelationship(ctx, playerID, characterID)
	if err != nil {
		return err
	}

	relationship.History = append(relationship.History, *interaction)
	relationship.UpdatedAt = time.Now()

	for _, effect := range interaction.Effects {
		switch effect.Type {
		case domain.EffectTypeAffection:
			relationship.Affection = math.Max(-100, math.Min(100, relationship.Affection+effect.Value))
		case domain.EffectTypeTrust:
			relationship.Trust = math.Max(0, math.Min(100, relationship.Trust+effect.Value))
		}
	}

	return s.UpdateRelationship(ctx, playerID, characterID, relationship)
}

func (s *StubRelationshipManager) CalculateCompatibility(ctx context.Context, player *domain.Player, character *domain.Character) (float64, error) {
	opennessDiff := math.Abs(player.Personality.Openness - character.Personality.Openness)
	extraversionDiff := math.Abs(player.Personality.Extraversion - character.Personality.Extraversion)
	agreeablenessDiff := math.Abs(player.Personality.Agreeableness - character.Personality.Agreeableness)

	totalDiff := opennessDiff + extraversionDiff + agreeablenessDiff
	compatibility := math.Max(0, 100-totalDiff)

	return compatibility, nil
}

func (s *StubRelationshipManager) GetRelationshipHistory(ctx context.Context, playerID, characterID string) ([]*domain.Interaction, error) {
	relationship, err := s.GetRelationship(ctx, playerID, characterID)
	if err != nil {
		return nil, err
	}

	interactions := make([]*domain.Interaction, len(relationship.History))
	for i := range relationship.History {
		interactions[i] = &relationship.History[i]
	}

	return interactions, nil
}

type StubEffectProcessor struct{}

func NewStubEffectProcessor() *StubEffectProcessor {
	return &StubEffectProcessor{}
}

func (s *StubEffectProcessor) ApplyEffect(ctx context.Context, effect *domain.Effect, gameState *domain.GameState) error {
	switch effect.Type {
	case domain.EffectTypeAffection:
		if relationship, exists := gameState.Player.Relationships[effect.Target]; exists {
			relationship.Affection = math.Max(-100, math.Min(100, relationship.Affection+effect.Value))
			gameState.Player.Relationships[effect.Target] = relationship
		}
	case domain.EffectTypeTrust:
		if relationship, exists := gameState.Player.Relationships[effect.Target]; exists {
			relationship.Trust = math.Max(0, math.Min(100, relationship.Trust+effect.Value))
			gameState.Player.Relationships[effect.Target] = relationship
		}
	case domain.EffectTypePopularity:
		gameState.Player.Stats.Popularity = math.Max(0, math.Min(100, gameState.Player.Stats.Popularity+effect.Value))
	case domain.EffectTypeConfidence:
		gameState.Player.Stats.Confidence = math.Max(0, math.Min(100, gameState.Player.Stats.Confidence+effect.Value))
	case domain.EffectTypeEnergy:
		gameState.Player.Stats.Energy = math.Max(0, math.Min(100, gameState.Player.Stats.Energy+effect.Value))
	case domain.EffectTypeMoney:
		gameState.Player.Stats.Money += int(effect.Value)
	}
	return nil
}

func (s *StubEffectProcessor) ApplyEffects(ctx context.Context, effects []*domain.Effect, gameState *domain.GameState) error {
	for _, effect := range effects {
		if err := s.ApplyEffect(ctx, effect, gameState); err != nil {
			return err
		}
	}
	return nil
}

func (s *StubEffectProcessor) ValidateEffect(ctx context.Context, effect *domain.Effect) error {
	return nil
}

func (s *StubEffectProcessor) ReverseEffect(ctx context.Context, effect *domain.Effect, gameState *domain.GameState) error {
	reversedEffect := &domain.Effect{
		Type:        effect.Type,
		Target:      effect.Target,
		Value:       -effect.Value,
		Description: "Reversed: " + effect.Description,
	}
	return s.ApplyEffect(ctx, reversedEffect, gameState)
}

type StubRequirementChecker struct{}

func NewStubRequirementChecker() *StubRequirementChecker {
	return &StubRequirementChecker{}
}

func (s *StubRequirementChecker) CheckRequirement(ctx context.Context, requirement *domain.Requirement, gameState *domain.GameState) (bool, error) {
	switch requirement.Type {
	case domain.RequirementTypeAffection:
		if relationship, exists := gameState.Player.Relationships[requirement.Target]; exists {
			return s.compareValues(relationship.Affection, requirement.Value, requirement.Operator), nil
		}
		return false, nil
	case domain.RequirementTypeTrust:
		if relationship, exists := gameState.Player.Relationships[requirement.Target]; exists {
			return s.compareValues(relationship.Trust, requirement.Value, requirement.Operator), nil
		}
		return false, nil
	case domain.RequirementTypePopularity:
		return s.compareValues(gameState.Player.Stats.Popularity, requirement.Value, requirement.Operator), nil
	case domain.RequirementTypeConfidence:
		return s.compareValues(gameState.Player.Stats.Confidence, requirement.Value, requirement.Operator), nil
	case domain.RequirementTypeEnergy:
		return s.compareValues(gameState.Player.Stats.Energy, requirement.Value, requirement.Operator), nil
	case domain.RequirementTypeMoney:
		return s.compareValues(float64(gameState.Player.Stats.Money), requirement.Value, requirement.Operator), nil
	case domain.RequirementTypeDayNumber:
		return s.compareValues(float64(gameState.Player.Stats.DayNumber), requirement.Value, requirement.Operator), nil
	default:
		return false, fmt.Errorf("unknown requirement type: %s", requirement.Type)
	}
}

func (s *StubRequirementChecker) CheckRequirements(ctx context.Context, requirements []*domain.Requirement, gameState *domain.GameState) (bool, error) {
	for _, requirement := range requirements {
		met, err := s.CheckRequirement(ctx, requirement, gameState)
		if err != nil {
			return false, err
		}
		if !met {
			return false, nil
		}
	}
	return true, nil
}

func (s *StubRequirementChecker) GetFailedRequirements(ctx context.Context, requirements []*domain.Requirement, gameState *domain.GameState) ([]*domain.Requirement, error) {
	var failed []*domain.Requirement
	for _, requirement := range requirements {
		met, err := s.CheckRequirement(ctx, requirement, gameState)
		if err != nil {
			return nil, err
		}
		if !met {
			failed = append(failed, requirement)
		}
	}
	return failed, nil
}

func (s *StubRequirementChecker) compareValues(actual, expected float64, operator string) bool {
	switch operator {
	case "eq":
		return actual == expected
	case "gt":
		return actual > expected
	case "lt":
		return actual < expected
	case "gte":
		return actual >= expected
	case "lte":
		return actual <= expected
	default:
		return false
	}
}

type StubConfigProvider struct{}

func NewStubConfigProvider() *StubConfigProvider {
	return &StubConfigProvider{}
}

func (s *StubConfigProvider) GetGameConfig(ctx context.Context) (*domain.GameConfig, error) {
	return &domain.GameConfig{
		MaxDays:            30,
		StartingMoney:      1000,
		StartingEnergy:     100.0,
		StartingConfidence: 50.0,
		StartingPopularity: 25.0,
		MaxCharacters:      6,
		EliminationDay:     15,
		FinaleDay:          30,
	}, nil
}

func (s *StubConfigProvider) GetEventConfigs(ctx context.Context) ([]*domain.Event, error) {
	return []*domain.Event{
		{
			ID:          "event_1",
			Title:       "Welcome to the Villa",
			Description: "You arrive at the Love Island villa and meet your fellow contestants.",
			Type:        domain.EventTypeDrama,
			Choices: []domain.Choice{
				{
					ID:          "choice_1",
					Text:        "Be confident and introduce yourself",
					Description: "Show your personality and make a good first impression",
					Effects: []domain.Effect{
						{Type: domain.EffectTypeConfidence, Target: "player", Value: 10.0, Description: "Confident introduction"},
						{Type: domain.EffectTypePopularity, Target: "player", Value: 5.0, Description: "Good first impression"},
					},
				},
				{
					ID:          "choice_2",
					Text:        "Stay quiet and observe",
					Description: "Take time to understand the dynamics before getting involved",
					Effects: []domain.Effect{
						{Type: domain.EffectTypeConfidence, Target: "player", Value: -5.0, Description: "Quiet observation"},
					},
				},
			},
			IsActive: true,
		},
		{
			ID:          "event_2",
			Title:       "First Challenge",
			Description: "The villa's first challenge tests your teamwork and communication.",
			Type:        domain.EventTypeChallenge,
			Choices: []domain.Choice{
				{
					ID:          "choice_3",
					Text:        "Take charge and lead the team",
					Description: "Show leadership qualities",
					Effects: []domain.Effect{
						{Type: domain.EffectTypeConfidence, Target: "player", Value: 15.0, Description: "Leadership shown"},
						{Type: domain.EffectTypeEnergy, Target: "player", Value: -10.0, Description: "Leadership effort"},
					},
				},
				{
					ID:          "choice_4",
					Text:        "Support your partner",
					Description: "Work together and build trust",
					Effects: []domain.Effect{
						{Type: domain.EffectTypeTrust, Target: "partner", Value: 10.0, Description: "Teamwork builds trust"},
					},
				},
			},
			IsActive: true,
		},
	}, nil
}

func (s *StubConfigProvider) GetCharacterConfigs(ctx context.Context) ([]*domain.Character, error) {
	return []*domain.Character{
		{
			ID:   "char_1",
			Name: "Emma",
			Age:  24,
			Personality: domain.Personality{
				Openness:          70.0,
				Conscientiousness: 60.0,
				Extraversion:      80.0,
				Agreeableness:     75.0,
				Neuroticism:       30.0,
			},
			Appearance: domain.Appearance{
				Height:    165,
				Build:     "Athletic",
				HairColor: "Blonde",
				EyeColor:  "Blue",
				Style:     "Casual chic",
			},
			Stats: domain.CharacterStats{
				Popularity: 60.0,
				Energy:     85.0,
				Stress:     20.0,
			},
			IsAvailable: true,
		},
		{
			ID:   "char_2",
			Name: "James",
			Age:  26,
			Personality: domain.Personality{
				Openness:          50.0,
				Conscientiousness: 80.0,
				Extraversion:      60.0,
				Agreeableness:     65.0,
				Neuroticism:       40.0,
			},
			Appearance: domain.Appearance{
				Height:    180,
				Build:     "Muscular",
				HairColor: "Brown",
				EyeColor:  "Green",
				Style:     "Smart casual",
			},
			Stats: domain.CharacterStats{
				Popularity: 55.0,
				Energy:     75.0,
				Stress:     30.0,
			},
			IsAvailable: true,
		},
		{
			ID:   "char_3",
			Name: "Sophie",
			Age:  23,
			Personality: domain.Personality{
				Openness:          85.0,
				Conscientiousness: 40.0,
				Extraversion:      90.0,
				Agreeableness:     70.0,
				Neuroticism:       50.0,
			},
			Appearance: domain.Appearance{
				Height:    170,
				Build:     "Slim",
				HairColor: "Red",
				EyeColor:  "Hazel",
				Style:     "Bohemian",
			},
			Stats: domain.CharacterStats{
				Popularity: 70.0,
				Energy:     90.0,
				Stress:     15.0,
			},
			IsAvailable: true,
		},
	}, nil
}

func (s *StubConfigProvider) GetItemConfigs(ctx context.Context) ([]*domain.Item, error) {
	return []*domain.Item{
		{
			ID:          "item_1",
			Name:        "Rose",
			Description: "A beautiful red rose",
			Type:        domain.ItemTypeGift,
			Value:       50,
		},
		{
			ID:          "item_2",
			Name:        "Chocolate",
			Description: "Delicious chocolate box",
			Type:        domain.ItemTypeGift,
			Value:       30,
		},
	}, nil
}
