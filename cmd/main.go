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

	// Create and display a person using our colors
	person := createExamplePerson()
	fmt.Printf("Person Created: %s\n",
		person.Name)
}

func initialize() {
	welcome()
	getInputs()
}

func fail()                   { fmt.Println("❌ You have failed to load ❌"); os.Exit(1) }
func success(nametype string) { fmt.Printf("✔ You have successfully loaded %s ✔ ", nametype) }
func welcome()                { welcomeScreen(); fmt.Println(CLEAR_SCREEN) }
func welcomeScreen()          { fmt.Println(BlueBackground + "Welcome to the game") }

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
