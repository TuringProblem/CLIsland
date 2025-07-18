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
	time.Sleep(time.Duration(2) * time.Second)
	TUIPrint(CLEAR_SCREEN + HOME)

	time.Sleep(time.Duration(2) * time.Second)

	loveMenu()
}

func loveMenu() {
	TUIPrint(CLEAR_SCREEN + HOME)
	myPrompts = addPromptAndSort(getSectionPrompts["main_menu"])
	for _, v := range myPrompts {
		TUIPrint(v)
	}
	handLoveMenuInput()

}

func handLoveMenuInput() {
	var input int
	TUIPrint("Please select an option:")
	myValue, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	switch myValue {
	case 1:
		characterBuild()
		break
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
	// build character by each part

	// sex
	mySex := buildSex()
	fmt.Println(mySex)

	// name

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
