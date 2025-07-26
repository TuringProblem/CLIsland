package main

func setSexForBuild(input int) (Sex, string) {
	if input == 1 {
		return Male, "Male"
	} else {
		return Female, "Female"
	}
}

func selectName(sex Sex) string {
	name := promptName()
	if sex == Male {
		TUIPrint(generateRandomNameFromFile(sex))
	} else {
		TUIPrint(generateRandomNameFromFile(sex))
	}
	return name
}
