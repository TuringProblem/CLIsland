package main

import (
	"fmt"
	"os"
)

func fail()                  { fmt.Println("❌ You have failed to load ❌"); os.Exit(1) }
func iGotATeeeexxxt() string { return "📱 I GOT A TEEEEEEEEEEXT!!!!!\n" }
func options() {
	printTag()
	fmt.Println("Please select an option:")
	fmt.Println("1. Start normal")
	fmt.Println("2. Start dev mode")
}
func printTag() string        { return "Deveoped by: @Override" }
func success(nametype string) { fmt.Printf("✔ You have successfully loaded %s ✔ ", nametype) }
