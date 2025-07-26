package utils

import (
	"testing"
)

func TestGenerateCharacters(t *testing.T) {
	maleCharacters := GenerateCharacters([]string{}, "male")
	if len(maleCharacters) != 11 {
		t.Errorf("Expected 11 characters for male player, got %d", len(maleCharacters))
	}

	femaleCharacters := GenerateCharacters([]string{}, "female")
	if len(femaleCharacters) != 11 {
		t.Errorf("Expected 11 characters for female player, got %d", len(femaleCharacters))
	}

	for i, name := range maleCharacters {
		if name == "" {
			t.Errorf("Character %d has empty name", i)
		}
	}

	for i, name := range femaleCharacters {
		if name == "" {
			t.Errorf("Character %d has empty name", i)
		}
	}
}

func TestGenerateRandomNameFromFile(t *testing.T) {
	maleName := GenerateRandomNameFromFile("male")
	if maleName == "" {
		t.Error("Generated male name is empty")
	}
	femaleName := GenerateRandomNameFromFile("female")
	if femaleName == "" {
		t.Error("Generated female name is empty")
	}
}
