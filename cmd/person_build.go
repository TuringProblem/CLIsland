package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"math/rand"
)


func buildName(sex Sex) string {
	printNameList(sex)
	myName := promptName()
	return myName
}

func promptName() string {
	TUIPrint(getSectionPrompts["character"]["name"])
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

// TODO: Clean this up, I don't like this redundancy - find a better way to understand go lol

func buildNameExample() string {
	var name string
	TUIPrint(getSectionPrompts["character"]["name"])
	fmt.Scanln(&name)
	return name
}

func buildSex() Sex {
	var sex Sex
	TUIPrint(getSectionPrompts["character"]["sex"])
	sex = sexConversion(&sex)
	return sex
}

func buildAge() int {
	var age int
	TUIPrint(getSectionPrompts["character"]["age"])
	fmt.Scanln(&age)
	return age
}

func buildHeight() float64 {
	var height float64
	TUIPrint(getSectionPrompts["character"]["height"])
	fmt.Scanln(&height)
	return height
}

func buildWeight() float64 {
	var weight float64
	TUIPrint(getSectionPrompts["character"]["weight"])
	fmt.Scanln(&weight)
	return weight
}

func sexConversion(sex *Sex) Sex {
	var input int
	input = intConversion(input)
	if input == 1 {
		*sex = Male
		return *sex
	} else {
		*sex = Female
		return *sex
	}
}

func intConversion(input int) int {
	response, err := fmt.Scan(&input)
	if err == nil {
		fmt.Errorf("Seems to be an error with %d", response)
	}
	return input
}




func generateRandomNameFromFile(sex Sex) string {
	if sex == Male {
		content, err := os.ReadFile("./data/names/boys.txt")
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), "\n")
		return lines[rand.Intn(len(lines))]
	} else {
		content, err := os.ReadFile("./data/names/girls.txt")
		if err != nil {
			log.Fatal(err)
		}
		lines := strings.Split(string(content), "\n")
		return lines[rand.Intn(len(lines))]
	}
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
