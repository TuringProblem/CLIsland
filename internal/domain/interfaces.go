package domain

import (
	"context"
)

type GameState struct {
	Player       *Player               `json:"player"`
	Characters   map[string]*Character `json:"characters"`
	Events       map[string]*Event     `json:"events"`
	CurrentEvent *Event                `json:"current_event"`
	GameDay      int                   `json:"game_day"`
	IsGameOver   bool                  `json:"is_game_over"`
	Winner       string                `json:"winner,omitempty"`
}

type GameEngine interface {
	StartGame(ctx context.Context, playerName string) (*GameState, error)
	ProcessChoice(ctx context.Context, choiceID string) (*GameState, error)
	AdvanceDay(ctx context.Context) (*GameState, error)
	EndGame(ctx context.Context) error

	GetCurrentState(ctx context.Context) (*GameState, error)
	SaveGame(ctx context.Context) error
	LoadGame(ctx context.Context) (*GameState, error)

	GetAvailableEvents(ctx context.Context) ([]*Event, error)
	TriggerEvent(ctx context.Context, eventID string) (*GameState, error)

	GetAvailableCharacters(ctx context.Context) ([]*Character, error)
	InteractWithCharacter(ctx context.Context, characterID string, interactionType InteractionType) (*GameState, error)
}

type EventManager interface {
	CreateEvent(ctx context.Context, event *Event) error
	GetEvent(ctx context.Context, eventID string) (*Event, error)
	UpdateEvent(ctx context.Context, event *Event) error
	DeleteEvent(ctx context.Context, eventID string) error
	ListEvents(ctx context.Context, eventType EventType) ([]*Event, error)
	ValidateEvent(ctx context.Context, event *Event) error
}

type CharacterManager interface {
	CreateCharacter(ctx context.Context, character *Character) error
	GetCharacter(ctx context.Context, characterID string) (*Character, error)
	UpdateCharacter(ctx context.Context, character *Character) error
	DeleteCharacter(ctx context.Context, characterID string) error
	ListCharacters(ctx context.Context) ([]*Character, error)
	UpdateCharacterStats(ctx context.Context, characterID string, stats CharacterStats) error
}

type RelationshipManager interface {
	GetRelationship(ctx context.Context, playerID, characterID string) (*Relationship, error)
	UpdateRelationship(ctx context.Context, playerID, characterID string, relationship *Relationship) error
	AddInteraction(ctx context.Context, playerID, characterID string, interaction *Interaction) error
	CalculateCompatibility(ctx context.Context, player *Player, character *Character) (float64, error)
	GetRelationshipHistory(ctx context.Context, playerID, characterID string) ([]*Interaction, error)
}

type EffectProcessor interface {
	ApplyEffect(ctx context.Context, effect *Effect, gameState *GameState) error
	ApplyEffects(ctx context.Context, effects []*Effect, gameState *GameState) error
	ValidateEffect(ctx context.Context, effect *Effect) error
	ReverseEffect(ctx context.Context, effect *Effect, gameState *GameState) error
}

type RequirementChecker interface {
	CheckRequirement(ctx context.Context, requirement *Requirement, gameState *GameState) (bool, error)
	CheckRequirements(ctx context.Context, requirements []*Requirement, gameState *GameState) (bool, error)
	GetFailedRequirements(ctx context.Context, requirements []*Requirement, gameState *GameState) ([]*Requirement, error)
}

type StateRepository interface {
	Save(ctx context.Context, gameState *GameState) error
	Load(ctx context.Context) (*GameState, error)
	Delete(ctx context.Context) error
	Exists(ctx context.Context) (bool, error)
}

type EventRepository interface {
	Save(ctx context.Context, event *Event) error
	GetByID(ctx context.Context, eventID string) (*Event, error)
	GetByType(ctx context.Context, eventType EventType) ([]*Event, error)
	GetAll(ctx context.Context) ([]*Event, error)
	Delete(ctx context.Context, eventID string) error
}

type CharacterRepository interface {
	Save(ctx context.Context, character *Character) error
	GetByID(ctx context.Context, characterID string) (*Character, error)
	GetAll(ctx context.Context) ([]*Character, error)
	Delete(ctx context.Context, characterID string) error
	Update(ctx context.Context, character *Character) error
}

type ConfigProvider interface {
	GetGameConfig(ctx context.Context) (*GameConfig, error)
	GetEventConfigs(ctx context.Context) ([]*Event, error)
	GetCharacterConfigs(ctx context.Context) ([]*Character, error)
	GetItemConfigs(ctx context.Context) ([]*Item, error)
}

type GameConfig struct {
	MaxDays            int     `json:"max_days"`
	StartingMoney      int     `json:"starting_money"`
	StartingEnergy     float64 `json:"starting_energy"`
	StartingConfidence float64 `json:"starting_confidence"`
	StartingPopularity float64 `json:"starting_popularity"`
	MaxCharacters      int     `json:"max_characters"`
	EliminationDay     int     `json:"elimination_day"`
	FinaleDay          int     `json:"finale_day"`
}
