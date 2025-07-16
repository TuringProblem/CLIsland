package main

func setSexForBuild(input int) (Sex, string) {
	if input == 1 {
		return Male, "Male"
	} else {
		return Female, "Female"
	}
}

func selectName(sex Sex) string {
	promptName()
	if sex == Male {
		return "Chris"
	} else {
		return "Huda"
	}
}
