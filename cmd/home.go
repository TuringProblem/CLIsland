package main

func start(message string) {
	TUIPrint(message)
	characterBuild()
}

func characterBuild() {
	TUIPrint("")
	for k, v := range getSectionPrompts["character"] {
		if k == "sex" {
			handlePromptIterator(v)
		}
	}
}
func handlePromptIterator(prompt string) {
	TUIPrint(prompt)
	success("GREAT SUCESSSSSS")
}
