package domain

import (
	"time"
)

// Player represents the main character in the game
type Player struct {
	ID              string                  `json:"id"`
	Name            string                  `json:"name"`
	Age             int                     `json:"age"`
	Personality     Personality             `json:"personality"`
	Relationships   map[string]Relationship `json:"relationships"`
	Stats           PlayerStats             `json:"stats"`
	Inventory       []Item                  `json:"inventory"`
	CurrentLocation string                  `json:"current_location"`
	CreatedAt       time.Time               `json:"created_at"`
	UpdatedAt       time.Time               `json:"updated_at"`
}

// Character represents other contestants in the villa
type Character struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Age         int            `json:"age"`
	Personality Personality    `json:"personality"`
	Appearance  Appearance     `json:"appearance"`
	Stats       CharacterStats `json:"stats"`
	IsAvailable bool           `json:"is_available"`
	CreatedAt   time.Time      `json:"created_at"`
}

// Event represents a game event or challenge
type Event struct {
	ID           string        `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Type         EventType     `json:"type"`
	Choices      []Choice      `json:"choices"`
	Requirements []Requirement `json:"requirements"`
	Outcomes     []Outcome     `json:"outcomes"`
	Duration     time.Duration `json:"duration"`
	IsActive     bool          `json:"is_active"`
	CreatedAt    time.Time     `json:"created_at"`
}

// Choice represents a player decision point
type Choice struct {
	ID           string        `json:"id"`
	Text         string        `json:"text"`
	Description  string        `json:"description"`
	Effects      []Effect      `json:"effects"`
	Requirements []Requirement `json:"requirements"`
}

// Outcome represents the result of an event or choice
type Outcome struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
	NextEventID string   `json:"next_event_id,omitempty"`
}

// Effect represents a change to game state
type Effect struct {
	Type        EffectType `json:"type"`
	Target      string     `json:"target"`
	Value       float64    `json:"value"`
	Description string     `json:"description"`
}

// Relationship represents the connection between characters
type Relationship struct {
	CharacterID   string             `json:"character_id"`
	Affection     float64            `json:"affection"`     // -100 to 100
	Trust         float64            `json:"trust"`         // 0 to 100
	Compatibility float64            `json:"compatibility"` // 0 to 100
	Status        RelationshipStatus `json:"status"`
	History       []Interaction      `json:"history"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

// Interaction represents a single interaction between characters
type Interaction struct {
	Type        InteractionType `json:"type"`
	Description string          `json:"description"`
	Effects     []Effect        `json:"effects"`
	Timestamp   time.Time       `json:"timestamp"`
}

// Personality represents character traits
type Personality struct {
	Openness          float64 `json:"openness"`          // 0-100
	Conscientiousness float64 `json:"conscientiousness"` // 0-100
	Extraversion      float64 `json:"extraversion"`      // 0-100
	Agreeableness     float64 `json:"agreeableness"`     // 0-100
	Neuroticism       float64 `json:"neuroticism"`       // 0-100
}

// Appearance represents physical characteristics
type Appearance struct {
	Height    int    `json:"height"`
	Build     string `json:"build"`
	HairColor string `json:"hair_color"`
	EyeColor  string `json:"eye_color"`
	Style     string `json:"style"`
}

// PlayerStats represents player's current stats
type PlayerStats struct {
	Popularity float64 `json:"popularity"` // 0-100
	Confidence float64 `json:"confidence"` // 0-100
	Energy     float64 `json:"energy"`     // 0-100
	Money      int     `json:"money"`
	DayNumber  int     `json:"day_number"`
}

// CharacterStats represents character's current stats
type CharacterStats struct {
	Popularity float64 `json:"popularity"` // 0-100
	Energy     float64 `json:"energy"`     // 0-100
	Stress     float64 `json:"stress"`     // 0-100
}

// Item represents collectible items
type Item struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        ItemType `json:"type"`
	Value       int      `json:"value"`
}

// Requirement represents conditions that must be met
type Requirement struct {
	Type     RequirementType `json:"type"`
	Target   string          `json:"target"`
	Value    float64         `json:"value"`
	Operator string          `json:"operator"` // "eq", "gt", "lt", "gte", "lte"
}

// Enums
type EventType string

const (
	EventTypeChallenge   EventType = "challenge"
	EventTypeDate        EventType = "date"
	EventTypeElimination EventType = "elimination"
	EventTypeDrama       EventType = "drama"
	EventTypeRecoupling  EventType = "recoupling"
)

type EffectType string

const (
	EffectTypeAffection  EffectType = "affection"
	EffectTypeTrust      EffectType = "trust"
	EffectTypePopularity EffectType = "popularity"
	EffectTypeConfidence EffectType = "confidence"
	EffectTypeEnergy     EffectType = "energy"
	EffectTypeMoney      EffectType = "money"
	EffectTypeItem       EffectType = "item"
)

type RelationshipStatus string

const (
	RelationshipStatusSingle    RelationshipStatus = "single"
	RelationshipStatusCoupled   RelationshipStatus = "coupled"
	RelationshipStatusExclusive RelationshipStatus = "exclusive"
	RelationshipStatusMarried   RelationshipStatus = "married"
)

type InteractionType string

const (
	InteractionTypeConversation InteractionType = "conversation"
	InteractionTypeDate         InteractionType = "date"
	InteractionTypeChallenge    InteractionType = "challenge"
	InteractionTypeGift         InteractionType = "gift"
	InteractionTypeArgument     InteractionType = "argument"
)

type ItemType string

const (
	ItemTypeClothing   ItemType = "clothing"
	ItemTypeAccessory  ItemType = "accessory"
	ItemTypeGift       ItemType = "gift"
	ItemTypeConsumable ItemType = "consumable"
)

type RequirementType string

const (
	RequirementTypeAffection  RequirementType = "affection"
	RequirementTypeTrust      RequirementType = "trust"
	RequirementTypePopularity RequirementType = "popularity"
	RequirementTypeConfidence RequirementType = "confidence"
	RequirementTypeEnergy     RequirementType = "energy"
	RequirementTypeMoney      RequirementType = "money"
	RequirementTypeItem       RequirementType = "item"
	RequirementTypeDayNumber  RequirementType = "day_number"
)
