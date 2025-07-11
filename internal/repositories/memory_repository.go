package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/TuringProblem/CLIsland/internal/domain"
)

type MemoryStateRepository struct {
	gameState *domain.GameState
	mu        sync.RWMutex
}

func NewMemoryStateRepository() *MemoryStateRepository {
	return &MemoryStateRepository{}
}

func (m *MemoryStateRepository) Save(ctx context.Context, gameState *domain.GameState) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gameState = gameState
	return nil
}

func (m *MemoryStateRepository) Load(ctx context.Context) (*domain.GameState, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.gameState == nil {
		return nil, fmt.Errorf("no game state found")
	}
	return m.gameState, nil
}

func (m *MemoryStateRepository) Delete(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.gameState = nil
	return nil
}

func (m *MemoryStateRepository) Exists(ctx context.Context) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.gameState != nil, nil
}

type MemoryEventRepository struct {
	events map[string]*domain.Event
	mu     sync.RWMutex
}

func NewMemoryEventRepository() *MemoryEventRepository {
	return &MemoryEventRepository{
		events: make(map[string]*domain.Event),
	}
}

func (m *MemoryEventRepository) Save(ctx context.Context, event *domain.Event) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.events[event.ID] = event
	return nil
}

func (m *MemoryEventRepository) GetByID(ctx context.Context, eventID string) (*domain.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	event, exists := m.events[eventID]
	if !exists {
		return nil, fmt.Errorf("event with ID %s not found", eventID)
	}
	return event, nil
}

func (m *MemoryEventRepository) GetByType(ctx context.Context, eventType domain.EventType) ([]*domain.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var events []*domain.Event
	for _, event := range m.events {
		if event.Type == eventType {
			events = append(events, event)
		}
	}
	return events, nil
}

func (m *MemoryEventRepository) GetAll(ctx context.Context) ([]*domain.Event, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var events []*domain.Event
	for _, event := range m.events {
		events = append(events, event)
	}
	return events, nil
}

func (m *MemoryEventRepository) Delete(ctx context.Context, eventID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.events, eventID)
	return nil
}

type MemoryCharacterRepository struct {
	characters map[string]*domain.Character
	mu         sync.RWMutex
}

func NewMemoryCharacterRepository() *MemoryCharacterRepository {
	return &MemoryCharacterRepository{
		characters: make(map[string]*domain.Character),
	}
}

func (m *MemoryCharacterRepository) Save(ctx context.Context, character *domain.Character) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.characters[character.ID] = character
	return nil
}

func (m *MemoryCharacterRepository) GetByID(ctx context.Context, characterID string) (*domain.Character, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	character, exists := m.characters[characterID]
	if !exists {
		return nil, fmt.Errorf("character with ID %s not found", characterID)
	}
	return character, nil
}

func (m *MemoryCharacterRepository) GetAll(ctx context.Context) ([]*domain.Character, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var characters []*domain.Character
	for _, character := range m.characters {
		characters = append(characters, character)
	}
	return characters, nil
}

func (m *MemoryCharacterRepository) Delete(ctx context.Context, characterID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.characters, characterID)
	return nil
}

func (m *MemoryCharacterRepository) Update(ctx context.Context, character *domain.Character) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.characters[character.ID] = character
	return nil
}
