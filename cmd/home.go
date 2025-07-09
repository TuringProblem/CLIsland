package main

import (
	"regexp"
)

func start(message string) {
	TUIPrint(message)
	characterBuild()
}

func characterBuild() {
	TUIPrint("")
	for _, prompt := range getSectionPrompts["character"] {
		handlePromptIterator(prompt)
	}
}
func handlePromptIterator(prompt string) {
	var sex string = "(sex)"
	match, _ := regexp.MatchString(sex, prompt)
	if match {
		TUIPrint(prompt)
		success("GREAT SUCESSSSSS")
	}
}
