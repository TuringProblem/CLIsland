package main

import (
	"fmt"
)

var WELCOME string = "Welcome to Love Island!"
var iGOTATEEEEEXT string = "📱 I GOT A TEEEEEEEEEEXT!!!!!\n"
var LOCAL_STRING string = "🏠"
var DEV_MODE_STRING string = "💻"
var PROD_STRING string = "🎬"

func options() {
	printTag()
	fmt.Println("Please select an option:")
	fmt.Println("1. Start normal")
	fmt.Println("2. Start dev mode")
}
func setMode(mode int) {
	switch mode {
	case 1:
		TUIPrint(LOCAL + LOCAL_STRING)
	case 2:
		TUIPrint(DEV_MODE + DEV_MODE_STRING)
	default:
		TUIPrint(greenMessage(PROD + PROD_STRING))
	}
}
func TUIPrint(message string) { fmt.Println(message) } // THIS IS THE MAIN PRINTING FUNCTION (JUST A WRAPPER AYO)
func greenMessage(message string) string {
	return GreenBackground + message + ResetBackground
}
func getTag() string          { return GreenBackground + "Deveoped by: @Override" + ResetBackground }
func printTag()               { TUIPrint(getTag()) }
func success(nametype string) { TUIPrint(GreenBackground + "✔ " + nametype + ResetBackground) }
func fail()                   { TUIPrint(RedBackground + "✘ FAIL" + ResetBackground) }
