package main

import (
	"fmt"
	"math/rand"
	"os"
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
	printTag()
	fmt.Println("Please select an option:")
	fmt.Println("1. Start normal")
	fmt.Println("2. Start dev mode")
	var input int
	fmt.Scan(&input)
	return input
}

func fail()                   { fmt.Println("❌ You have failed to load ❌"); os.Exit(1) }
func success(nametype string) { fmt.Printf("✔ You have successfully loaded %s ✔ ", nametype) }
func welcome() {
	fmt.Println(CLEAR_SCREEN)
	which(design())

}
func welcomeScreen(person Person) {
	fmt.Println(BlueBackground + "Welcome to the game" + person.Name + ResetBackground)
}
func which(option int) string {
	switch option {
	case 1:
		return "1"
	case 2:
		return "2"
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
