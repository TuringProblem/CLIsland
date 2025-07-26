package main

import (
	"fmt"

	"github.com/TuringProblem/CLIsland/utils"
)

func main() {
	fmt.Println("=== Love Island Person Generation Demo ===\n")

	// Generate people for a male player
	fmt.Println("🏝️  Characters for Male Player:")
	malePlayerPeople := utils.GeneratePersonList("male")
	for i, person := range malePlayerPeople {
		fmt.Printf("%d. %s (%s, %d years old, %s)\n",
			i+1, person.Name, person.Sex, person.Age, person.GetHeightInFeet())
		fmt.Printf("   Weight: %.1f kg (%.1f lbs)\n", person.GetWeightInKg(), person.GetWeightInLbs())
		fmt.Printf("   Interests: ")
		for j, interest := range person.GetInterests() {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%d/10)", interest, person.GetInterestWeight(interest))
		}
		fmt.Println("\n")
	}

	fmt.Println("🏝️  Characters for Female Player:")
	femalePlayerPeople := utils.GeneratePersonList("female")
	for i, person := range femalePlayerPeople {
		fmt.Printf("%d. %s (%s, %d years old, %s)\n",
			i+1, person.Name, person.Sex, person.Age, person.GetHeightInFeet())
		fmt.Printf("   Weight: %.1f kg (%.1f lbs)\n", person.GetWeightInKg(), person.GetWeightInLbs())
		fmt.Printf("   Interests: ")
		for j, interest := range person.GetInterests() {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (%d/10)", interest, person.GetInterestWeight(interest))
		}
		fmt.Println("\n")
	}

	// Show some statistics
	fmt.Println("=== Statistics ===")
	fmt.Printf("Male player gets: %d characters\n", len(malePlayerPeople))
	fmt.Printf("Female player gets: %d characters\n", len(femalePlayerPeople))

	// Count sexes for male player
	maleCount := 0
	femaleCount := 0
	for _, person := range malePlayerPeople {
		if person.Sex == "male" {
			maleCount++
		} else {
			femaleCount++
		}
	}
	fmt.Printf("Male player breakdown: %d males, %d females\n", maleCount, femaleCount)

	// Count sexes for female player
	maleCount = 0
	femaleCount = 0
	for _, person := range femalePlayerPeople {
		if person.Sex == "male" {
			maleCount++
		} else {
			femaleCount++
		}
	}
	fmt.Printf("Female player breakdown: %d males, %d females\n", maleCount, femaleCount)
}
