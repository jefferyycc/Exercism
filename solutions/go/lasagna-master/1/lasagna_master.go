package lasagnamaster

func PreparationTime(layers []string, prepTime int) int {
	if prepTime == 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

func Quantities(ingredients []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, ingred := range ingredients {
		switch ingred {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	seen := make(map[string]bool, len(myList))
	for _, ingredient := range myList {
		seen[ingredient] = true
	}

	for _, f := range friendsList {
		if !seen[f] {
			myList[len(myList)-1] = f
			seen[f] = true
		}
	}
	return myList
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	factor := float64(portions) / 2.0
	scaled := make([]float64, len(quantities))
	for i, q := range quantities {
		scaled[i] = q * factor
	}
	return scaled
}
