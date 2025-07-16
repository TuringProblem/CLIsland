package main

import "fmt"

var getSectionPrompts map[string]map[string]string = map[string]map[string]string{
	"character": {
		"name":   "What is your name?",
		"age":    "What is your age?",
		"height": "What is your height?",
		"weight": "What is your weight?",
		"sex":    "What is your sex?",
	},
	"main_menu": {
		"header":       "Welcome to Love Island!",
		"option_one":   "[1]: Start",
		"option_two":   "[2]: Settings",
		"option_three": "[3]: Exit",
	},
	"settings": {
		"header": "Settings",
	},
	"intro": {
		"header": "Welcome to Love Island!",
	},
}

func namePrompt() string {
	var name string
	TUIPrint("What is your name?")
	fmt.Scanln(&name)
	return name
}

func agePrompt() int {
	var age int
	TUIPrint("What is your age?")
	fmt.Scanln(&age)
	return age
}
