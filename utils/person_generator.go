package utils

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name      string
	Age       int
	Height    float64 // inches
	Weight    float64
	Sex       string
	Interests Interests
}

type Interest string

const (
	Music    Interest = "Music"
	Sports   Interest = "Sports"
	Reading  Interest = "Reading"
	Writing  Interest = "Writing"
	Coding   Interest = "Coding"
	Art      Interest = "Art"
	Travel   Interest = "Travel"
	Swimming Interest = "Swimming"
	Gaming   Interest = "Gaming"
	Lifting  Interest = "Lifting"
	Cooking  Interest = "Cooking"
	Cleaning Interest = "Cleaning"
	Shopping Interest = "Shopping"
	Partying Interest = "Partying"
	Sleeping Interest = "Sleeping"
)

type Interests struct {
	InterestType map[Interest]int // returns the weight of the interest in that category
}

func generateRandomAge() int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(18) + 18 // 18-35 years old
}

func generateRandomHeight(sex string) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	if sex == "male" {
		return float64(rand.Intn(14) + 65)
	} else {
		return float64(rand.Intn(13) + 60)
	}
}

func generateRandomWeight(height float64, sex string) float64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	minBMI := 18.5
	maxBMI := 25.0

	heightMeters := height * 0.0254

	minWeight := minBMI * heightMeters * heightMeters
	maxWeight := maxBMI * heightMeters * heightMeters

	weightRange := maxWeight - minWeight
	randomWeight := minWeight + (rand.Float64() * weightRange)

	return float64(int(randomWeight*10)) / 10
}

func generateRandomInterests() Interests {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	allInterests := []Interest{
		Music, Sports, Reading, Writing, Coding, Art, Travel, Swimming,
		Gaming, Lifting, Cooking, Cleaning, Shopping, Partying, Sleeping,
	}

	interests := Interests{
		InterestType: make(map[Interest]int),
	}

	// Select 3-6 random interests
	numInterests := rand.Intn(4) + 3 // 3-6 interests

	// Shuffle interests and take the first numInterests
	shuffled := make([]Interest, len(allInterests))
	copy(shuffled, allInterests)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// Assign random weights (1-10) to selected interests
	for i := 0; i < numInterests; i++ {
		weight := rand.Intn(10) + 1 // 1-10
		interests.InterestType[shuffled[i]] = weight
	}

	return interests
}

// GeneratePerson creates a complete Person object with random attributes
func GeneratePerson(sex string) Person {
	name := GenerateRandomNameFromFile(sex)
	age := generateRandomAge()
	height := generateRandomHeight(sex)
	weight := generateRandomWeight(height, sex)
	interests := generateRandomInterests()

	return Person{
		Name:      name,
		Age:       age,
		Height:    height,
		Weight:    weight,
		Sex:       sex,
		Interests: interests,
	}
}

// GeneratePersonList creates a list of Person objects based on the player's sex
// If playerSex is "male", generates 5 boys and 6 girls
// If playerSex is "female", generates 6 boys and 5 girls
func GeneratePersonList(playerSex string) []Person {
	var people []Person

	if playerSex == "male" {
		// Generate 5 boys and 6 girls for male player
		for i := 0; i < 5; i++ {
			people = append(people, GeneratePerson("male"))
		}
		for i := 0; i < 6; i++ {
			people = append(people, GeneratePerson("female"))
		}
	} else {
		// Generate 6 boys and 5 girls for female player
		for i := 0; i < 6; i++ {
			people = append(people, GeneratePerson("male"))
		}
		for i := 0; i < 5; i++ {
			people = append(people, GeneratePerson("female"))
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(people), func(i, j int) {
		people[i], people[j] = people[j], people[i]
	})

	return people
}

// GetInterests returns a slice of interests for a person
func (p *Person) GetInterests() []Interest {
	var interests []Interest
	for k := range p.Interests.InterestType {
		interests = append(interests, k)
	}
	return interests
}

// GetInterestWeight returns the weight of a specific interest
func (p *Person) GetInterestWeight(interest Interest) int {
	return p.Interests.InterestType[interest]
}

// HasInterest checks if a person has a specific interest
func (p *Person) HasInterest(interest Interest) bool {
	_, exists := p.Interests.InterestType[interest]
	return exists
}

// GetHeightInFeet returns height in feet and inches format
func (p *Person) GetHeightInFeet() string {
	feet := int(p.Height) / 12
	inches := int(p.Height) % 12
	return fmt.Sprintf("%d'%d\"", feet, inches)
}

// GetHeightInCm returns height in centimeters
func (p *Person) GetHeightInCm() float64 {
	return p.Height * 2.54
}

// GetWeightInKg returns weight in kilograms
func (p *Person) GetWeightInKg() float64 {
	return p.Weight
}

// GetWeightInLbs returns weight in pounds
func (p *Person) GetWeightInLbs() float64 {
	return p.Weight * 2.20462
}
