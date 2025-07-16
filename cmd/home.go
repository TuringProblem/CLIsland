package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

var myPrompts []string

func start(person Person) {
	TUIPrint(iGOTATEEEEEXT)
	TUIPrint(person.Name)
	person.PrintPersonInterestsPretty()
	time.Sleep(time.Duration(3) * time.Second)
	TUIPrint(CLEAR_SCREEN + HOME)

	time.Sleep(time.Duration(3) * time.Second)

	loveMenu()
}

func loveMenu() {
	TUIPrint(CLEAR_SCREEN + HOME)
	myPrompts = addPromptAndSort(getSectionPrompts["main_menu"])
	for _, v := range myPrompts {
		TUIPrint(v)
	}
	var input int
	handLoveMenuInput(input)

}

func handLoveMenuInput(input int) {
	TUIPrint("Please select an option:")
	fmt.Scan(&input)
	switch input {
	case 1:
		characterBuild()
	case 2:
		//settingsBuild()
		TUIPrint(CLEAR_SCREEN + HOME)
		loveMenu()
	case 3:
		os.Exit(0)
	default:
		fail()
	}
}

func characterBuild() {
	TUIPrint(CLEAR_SCREEN + HOME)
	for k, v := range getSectionPrompts["character"] {
		if k == "sex" {
    myValue := handlePromptIterator(v)
	}
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

func addPromptAndSort(prompt map[string]string) []string {
	for _, v := range prompt {
		myPrompts = append(myPrompts, v)
	}
	sort.Slice(myPrompts, func(i, j int) bool {
		return myPrompts[i] < myPrompts[j]
	})
	return myPrompts
}
