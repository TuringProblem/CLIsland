package utils

import "fmt"

func ExampleUsage() {
	malePlayerCharacters := GenerateCharacters([]string{}, "male")
	fmt.Printf("Characters for male player: %v\n", malePlayerCharacters)
	fmt.Printf("Total characters: %d\n", len(malePlayerCharacters))

	femalePlayerCharacters := GenerateCharacters([]string{}, "female")
	fmt.Printf("Characters for female player: %v\n", femalePlayerCharacters)
	fmt.Printf("Total characters: %d\n", len(femalePlayerCharacters))
}
