package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	souce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			souce += 0.2
		}
	}
	return noodles, souce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaledQuantities := make([]float64, 0, len(quantities))
	multiplier := float64(portions) / 2.0
	for _, quntity := range quantities {
		scaledQuantities = append(scaledQuantities, quntity*multiplier)
	}

	return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
