package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func main() {
	printTag()
	initialize()
	TUIPrint("✔Hello, World")
}

func initialize() { welcome() }
func welcome()    { TUIPrint(CLEAR_SCREEN); which(versionResponse()) }

func versionResponse() int {
	options()
	var input int
	fmt.Scan(&input)
	return input
}

func welcomeScreen(person Person) {
	TUIPrint(BlueBackground + "Welcome to the game " + person.Name + ResetBackground)
}

func which(option int) {
	switch option {
	case 1:
		setMode(1)
		welcomeScreen(createExamplePerson())
	case 2:
		TUIPrint(CLEAR_SCREEN)
		setMode(2)
		start(iGotATeeeexxxt())
	default:
		fail()
	}
}

/**
func getInputs() {
	fmt.Println("Grabbing inputs...")
	var rand int = rand.Intn(10)
	if rand > 5 {
		time.Sleep(time.Duration(5) * time.Second)
		fail()
	} else {
		time.Sleep(time.Duration(2) * time.Second)
		success("inputs")
		time.Sleep(time.Duration(1) * time.Second)
	}
}
**/
