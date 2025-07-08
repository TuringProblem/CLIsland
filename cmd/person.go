package main

type Sex string

const (
	Male   Sex = "Male"
	Female Sex = "Female"
)

type Interest string

const (
	Music   Interest = "Music"
	Sports  Interest = "Sports"
	Reading Interest = "Reading"
	Writing Interest = "Writing"
	Coding  Interest = "Coding"
	Art     Interest = "Art"
	Travel  Interest = "Travel"
)

type Interests struct {
	InterestType map[Interest]int // returns the weight of the interest in that category
}

type Person struct {
	Name      string
	Age       int
	Height    float64 // inches
	Weight    float64
	sex       Sex
	Interests Interests
	Hobbies   []string
}

// Example usage function
func createExamplePerson() Person {
	myPerson := Person{
		Name:   "Andrew",
		Age:    30, //can't be below 18
		Height: 73, // TODO: Need to make a converter to go from inches to ft. :) && also to cm for UK brev
		Weight: 70,
		sex:    Male,
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
		Hobbies: []string{"Coding", "Travel"},
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
