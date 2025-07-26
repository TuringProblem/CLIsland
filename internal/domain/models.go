package domain

import (
	"time"
)

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

type Choice struct {
	ID           string        `json:"id"`
	Text         string        `json:"text"`
	Description  string        `json:"description"`
	Effects      []Effect      `json:"effects"`
	Requirements []Requirement `json:"requirements"`
}

type Outcome struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Effects     []Effect `json:"effects"`
	NextEventID string   `json:"next_event_id,omitempty"`
}

type Effect struct {
	Type        EffectType `json:"type"`
	Target      string     `json:"target"`
	Value       float64    `json:"value"`
	Description string     `json:"description"`
}

type Relationship struct {
	CharacterID   string             `json:"character_id"`
	Affection     float64            `json:"affection"`     // -100 to 100
	Trust         float64            `json:"trust"`         // 0 to 100
	Compatibility float64            `json:"compatibility"` // 0 to 100
	Status        RelationshipStatus `json:"status"`
	History       []Interaction      `json:"history"`
	UpdatedAt     time.Time          `json:"updated_at"`
}

type Interaction struct {
	Type        InteractionType `json:"type"`
	Description string          `json:"description"`
	Effects     []Effect        `json:"effects"`
	Timestamp   time.Time       `json:"timestamp"`
}

type Personality struct {
	Openness          float64 `json:"openness"`          // 0-100
	Conscientiousness float64 `json:"conscientiousness"` // 0-100
	Extraversion      float64 `json:"extraversion"`      // 0-100
	Agreeableness     float64 `json:"agreeableness"`     // 0-100
	Neuroticism       float64 `json:"neuroticism"`       // 0-100
}

type Appearance struct {
	Height    int    `json:"height"`
	Build     string `json:"build"`
	HairColor string `json:"hair_color"`
	EyeColor  string `json:"eye_color"`
	Style     string `json:"style"`
}

type PlayerStats struct {
	Popularity float64 `json:"popularity"` // 0-100
	Confidence float64 `json:"confidence"` // 0-100
	Energy     float64 `json:"energy"`     // 0-100
	Money      int     `json:"money"`
	DayNumber  int     `json:"day_number"`
}

type CharacterStats struct {
	Popularity float64 `json:"popularity"` // 0-100
	Energy     float64 `json:"energy"`     // 0-100
	Stress     float64 `json:"stress"`     // 0-100
}

type Item struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        ItemType `json:"type"`
	Value       int      `json:"value"`
}

type Requirement struct {
	Type     RequirementType `json:"type"`
	Target   string          `json:"target"`
	Value    float64         `json:"value"`
	Operator string          `json:"operator"` // "eq", "gt", "lt", "gte", "lte"
}

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
