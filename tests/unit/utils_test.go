package tests

import (
	"testing"

	"github.com/TuringProblem/CLIsland/utils"
)

func TestGenerateCharacters(t *testing.T) {
	maleCharacters := utils.GenerateCharacters([]string{}, "male")
	if len(maleCharacters) != 11 {
		t.Errorf("Expected 11 characters for male player, got %d", len(maleCharacters))
	}

	femaleCharacters := utils.GenerateCharacters([]string{}, "female")
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
	maleName := utils.GenerateRandomNameFromFile("male")
	if maleName == "" {
		t.Error("Generated male name is empty")
	}

	femaleName := utils.GenerateRandomNameFromFile("female")
	if femaleName == "" {
		t.Error("Generated female name is empty")
	}
}

func TestGeneratePerson(t *testing.T) {
	malePerson := utils.GeneratePerson("male")
	if malePerson.Name == "" {
		t.Error("Generated male person has empty name")
	}
	if malePerson.Age < 18 || malePerson.Age > 35 {
		t.Errorf("Generated male person age %d is outside valid range (18-35)", malePerson.Age)
	}
	if malePerson.Sex != "male" {
		t.Errorf("Generated male person has wrong sex: %s", malePerson.Sex)
	}
	if malePerson.Height < 65 || malePerson.Height > 78 {
		t.Errorf("Generated male person height %.1f is outside valid range (65-78 inches)", malePerson.Height)
	}
	if len(malePerson.GetInterests()) < 3 || len(malePerson.GetInterests()) > 6 {
		t.Errorf("Generated male person has %d interests, expected 3-6", len(malePerson.GetInterests()))
	}

	femalePerson := utils.GeneratePerson("female")
	if femalePerson.Name == "" {
		t.Error("Generated female person has empty name")
	}
	if femalePerson.Age < 18 || femalePerson.Age > 35 {
		t.Errorf("Generated female person age %d is outside valid range (18-35)", femalePerson.Age)
	}
	if femalePerson.Sex != "female" {
		t.Errorf("Generated female person has wrong sex: %s", femalePerson.Sex)
	}
	if femalePerson.Height < 60 || femalePerson.Height > 72 {
		t.Errorf("Generated female person height %.1f is outside valid range (60-72 inches)", femalePerson.Height)
	}
	if len(femalePerson.GetInterests()) < 3 || len(femalePerson.GetInterests()) > 6 {
		t.Errorf("Generated female person has %d interests, expected 3-6", len(femalePerson.GetInterests()))
	}
}

func TestGeneratePersonList(t *testing.T) {
	malePlayerPeople := utils.GeneratePersonList("male")
	if len(malePlayerPeople) != 11 {
		t.Errorf("Expected 11 people for male player, got %d", len(malePlayerPeople))
	}

	femalePlayerPeople := utils.GeneratePersonList("female")
	if len(femalePlayerPeople) != 11 {
		t.Errorf("Expected 11 people for female player, got %d", len(femalePlayerPeople))
	}

	for i, person := range malePlayerPeople {
		if person.Name == "" {
			t.Errorf("Person %d has empty name", i)
		}
		if person.Age < 18 || person.Age > 35 {
			t.Errorf("Person %d has invalid age: %d", i, person.Age)
		}
		if len(person.GetInterests()) < 3 || len(person.GetInterests()) > 6 {
			t.Errorf("Person %d has invalid number of interests: %d", i, len(person.GetInterests()))
		}
	}
}

func TestPersonMethods(t *testing.T) {
	person := utils.GeneratePerson("male")

	heightFeet := person.GetHeightInFeet()
	if heightFeet == "" {
		t.Error("GetHeightInFeet returned empty string")
	}

	heightCm := person.GetHeightInCm()
	if heightCm <= 0 {
		t.Error("GetHeightInCm returned invalid value")
	}

	weightKg := person.GetWeightInKg()
	if weightKg <= 0 {
		t.Error("GetWeightInKg returned invalid value")
	}

	weightLbs := person.GetWeightInLbs()
	if weightLbs <= 0 {
		t.Error("GetWeightInLbs returned invalid value")
	}

	interests := person.GetInterests()
	if len(interests) > 0 {
		firstInterest := interests[0]
		if !person.HasInterest(firstInterest) {
			t.Errorf("HasInterest returned false for interest %s that should exist", firstInterest)
		}
	}

	if len(interests) > 0 {
		firstInterest := interests[0]
		weight := person.GetInterestWeight(firstInterest)
		if weight < 1 || weight > 10 {
			t.Errorf("GetInterestWeight returned invalid weight %d for interest %s", weight, firstInterest)
		}
	}
}
