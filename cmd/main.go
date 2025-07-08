package main

import (
	"fmt"
	"math/rand"
	"time"
)

const CLEAR_SCREEN = "\033[2J"

func main() {
	initialize()
	fmt.Println("✔Hello, World")

}

func initialize() {
	welcome()
	time.Sleep(time.Duration(3) * time.Second)
	getInputs()
}

func design() int {
	options()
	var input int
	fmt.Scan(&input)
	return input
}
func welcome() {
	fmt.Println(CLEAR_SCREEN)
	answer := which(design())
	fmt.Println(answer)
}
func welcomeScreen(person Person) {
	fmt.Println(BlueBackground + "Welcome to the game" + person.Name + ResetBackground)
}
func which(option int) string {
	switch option {
	case 1:
		return "1"
	case 2:
		return iGotATeeeexxxt()
	default:
		return "You must enter either [1] or [2]"
	}
}
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
