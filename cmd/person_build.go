package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func buildName(sex Sex) string {
	printNameList(sex)
	myName := promptName()
	return myName
}

func promptName() string {
	fmt.Println("Please enter a name:")
	var name string
	fmt.Scanln(&name)
	return name
}
func heightWeightPrompt() (float64, float64) {
	var height, weight float64
	fmt.Println("Please enter your height in inches:")
	fmt.Scanln(&height)
	fmt.Println("Please enter your weight in pounds:")
	fmt.Scanln(&weight)
	return height, weight
}

func agePrompt() int {
	var age int
	TUIPrint("What is your age?")
	fmt.Scanln(&age)
	return age
}

func printNameList(sex Sex) {
	var content []byte
	var err error
	if sex == Male {
		content, err = os.ReadFile("data/boys.txt")
	} else {
		content, err = os.ReadFile("data/girls.txt")
	}
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
}
