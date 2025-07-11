package main

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
}
