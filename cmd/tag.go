package main

import (
	"fmt"
)

func options() {
	printTag()
	fmt.Println("Please select an option:")
	fmt.Println("1. Start normal")
	fmt.Println("2. Start dev mode")
}
func setMode(mode int) {
	switch mode {
	case 1:
		TUIPrint(LOCAL)
	case 2:
		TUIPrint(DEV_MODE)
	default:
		TUIPrint(PROD)
	}
}
func TUIPrint(message string) { fmt.Println(message) } // THIS IS THE MAIN PRINTING FUNCTION (JUST A WRAPPER AYO)
func getTag() string          { return GreenBackground + "Deveoped by: @Override" + ResetBackground }
func printTag()               { TUIPrint(getTag()) }
func iGotATeeeexxxt() string  { return "📱 I GOT A TEEEEEEEEEEXT!!!!!\n" }
func success(nametype string) { TUIPrint(GreenBackground + "✔ " + nametype + ResetBackground) }
func fail()                   { TUIPrint(RedBackground + "✘ FAIL" + ResetBackground) }
