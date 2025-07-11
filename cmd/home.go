package main

import (
	"fmt"
)

func start(message string) {
	TUIPrint(message)
	loveMenu()
	characterBuild()
}

func loveMenu() {
	TUIPrint(CLEAR_SCREEN + HOME)
	for _, v := range getSectionPrompts["main_menu"] {
		TUIPrint(v)
	}
}
func characterBuild() {
	TUIPrint(CLEAR_SCREEN)
	for k, v := range getSectionPrompts["character"] {
		if k == "sex" {
			handlePromptIterator(v)
		}
	}
	examplePerson := createExamplePerson()

	examplePerson.PrintPersonInterestsPretty()

}
func handlePromptIterator(prompt string) int {
	var response int
	TUIPrint(prompt)
	success("Please Select:\n[1] male [2] female")
	response, err := fmt.Scan(&response)
	if err == nil {
		fmt.Errorf("Seems to be an error with %d", response)
	}

	return response
}
