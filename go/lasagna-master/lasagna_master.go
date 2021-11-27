package lasagna

// func PreparationTime estimates how long preparation will take based on layers count.
func PreparationTime(layers []string, avrPrepTime int) int {
	if avrPrepTime <= 0 {
		avrPrepTime = 2
	}
	return len(layers) * avrPrepTime
}

// func Quantities computes the amounts of noodles and sauce needed.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0
	for _, v := range layers {
		switch v {
		case "noodles":
			noodles++
		case "sauce":
			sauce++
		}
	}
	return noodles * 50, float64(sauce) * 0.2
}

// func AddSecretIngredient add secret ingredient.
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// func ScaleRecipe scales the recipe.
func ScaleRecipe(quantites []float64, portions int) []float64 {
	r := make([]float64, len(quantites), len(quantites))
	for i, v := range quantites {
		r[i] = v * float64(portions) / 2
	}
	return r
}
