package main

import (
	"fmt"
	"math/rand"
)

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

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

type Person struct {
	Name      string
	Age       int
	Height    float64 // inches
	Weight    float64
	Sex       Sex
	Interests Interests
}

func createPerson() *Person {
	myPerson := Person{
		Name: buildName(getRandomSex()),
	}
	return &myPerson
}

func buildPerson() Person {
	p := Person{}
	p.Sex, _ = setSexForBuild(p.mySexAsInt())
	p.Name = selectName(p.Sex)
	p.Age = rand.Intn(100)
	p.Height, p.Weight = heightWeightPrompt()
	return p
}

func (p *Person) mySexAsInt() int {
	if p.Sex == Male {
		return 1
	} else {
		return 2
	}
}

func createRandomPerson() Person {
	myPerson := Person{
		Name:   "",
		Age:    rand.Intn(100),
		Height: rand.Float64() * 100,
		Weight: rand.Float64() * 100,
		Sex:    getRandomSex(),
		Interests: Interests{
			InterestType: map[Interest]int{
				Music:   rand.Intn(10),
				Sports:  rand.Intn(10),
				Reading: rand.Intn(10),
				Writing: rand.Intn(10),
				Coding:  rand.Intn(10),
				Art:     rand.Intn(10),
				Travel:  rand.Intn(10),
			},
		},
	}
	myPerson.Name = selectName(myPerson.Sex)
	return myPerson
}
func getRandomSex() Sex {
	randomValue := rand.Intn(2)
	if randomValue == 0 {
		return Male
	} else {
		return Female
	}
}

func createExamplePerson() Person {
	myPerson := Person{
		Name:   "Andrew",
		Age:    30, //can't be below 18
		Height: 73, // TODO: Need to make a converter to go from inches to ft. :) && also to cm for UK brev
		Weight: 70,
		Sex:    getRandomSex(),
		Interests: Interests{
			InterestType: map[Interest]int{
				Music:   5,
				Sports:  10,
				Reading: 3,
				Writing: 2,
				Coding:  1,
				Art:     1,
				Travel:  1,
			},
		},
	}

	return myPerson
}

func (p *Person) GetInterests() []Interest {
	var interests []Interest
	for k := range p.Interests.InterestType {
		interests = append(interests, k)
	}
	return interests
}
func (p *Person) PrintPersonInterestsPretty() {
	TUIPrint("[=== Interests ===]")
	for k, v := range p.Interests.InterestType {
		fmt.Printf("%s[ %s ]%s\n%s%s\n", UNDER, k, RES, printAsBlock(v), ResetBackground)
	}
}

// TODO: Figure out a more clever way to solve this -> feel like maybe a for loop... even a map
func printAsBlock(interestWeight int) string {
	switch interestWeight {
	case 1:
		return fmt.Sprintf("[ %s█%s | | | | | | | | | ]", Green, Reset)
	case 2:
		return fmt.Sprintf("[ %s█ █%s | | | | | | | | ]", Green, Reset)
	case 3:
		return fmt.Sprintf("[ %s█ █ █%s | | | | | | | ]", Green, Reset)
	case 4:
		return fmt.Sprintf("[ %s█ █ █ █%s | | | | | | ]", Green, Reset)
	case 5:
		return fmt.Sprintf("[ %s█ █ █ █ █%s | | | | | ]", Green, Reset)
	case 6:
		return fmt.Sprintf("[ %s█ █ █ █ █ █%s | | | | ]", Green, Reset)
	case 7:
		return fmt.Sprintf("[ %s█ █ █ █ █ █ █%s | | | ]", Green, Reset)
	case 8:
		return fmt.Sprintf("[ %s█ █ █ █ █ █ █ █%s | | ]", Green, Reset)
	case 9:
		return fmt.Sprintf("[ %s█ █ █ █ █ █ █ █ █%s | ]", Green, Reset)
	case 10:
		return fmt.Sprintf("[ %s█ █ █ █ █ █ █ █ █ █%s ]", Green, Reset)
	default:
		return "[ | | | | | | | | | | ]"
	}
}
