package main

var getSectionPrompts map[string]map[string]string = map[string]map[string]string{
	"character": {
		"name":   "What is your name?\nPlease enter: ",
		"age":    "What is your age?\nPlease enter: ",
		"height": "What is your height?\nPlease enter: ",
		"weight": "What is your weight\nPlease enter: ",
		"sex":    "What is your sex?\n[1] male [2] female: ",
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
