package utils

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GenerateRandomNameFromFile(sex string) string {
	var filename string
	if sex == "male" {
		filename = "data/names/boys.txt"
	} else {
		filename = "data/names/girls.txt"
	}

	file, err := os.Open(filename)
	if err != nil {
		if sex == "male" {
			return "Alex"
		}
		return "Emma"
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if name != "" {
			names = append(names, name)
		}
	}

	if len(names) == 0 {
		if sex == "male" {
			return "Alex"
		}
		return "Emma"
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	return names[rand.Intn(len(names))]
}

func ShuffleStrings(slice []string) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// GenerateCharacters generates a list of character names based on the player's sex
// If playerSex is "male", generates 5 boys and 6 girls
// If playerSex is "female", generates 6 boys and 5 girls
func GenerateCharacters(characters []string, playerSex string) []string {
	var result []string

	if playerSex == "male" {
		for i := 0; i < 5; i++ {
			result = append(result, GenerateRandomNameFromFile("male"))
		}
		for i := 0; i < 6; i++ {
			result = append(result, GenerateRandomNameFromFile("female"))
		}
	} else {
		for i := 0; i < 6; i++ {
			result = append(result, GenerateRandomNameFromFile("male"))
		}
		for i := 0; i < 5; i++ {
			result = append(result, GenerateRandomNameFromFile("female"))
		}
	}

	ShuffleStrings(result)

	return result
}
