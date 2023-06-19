package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, layerPrepTime int) int {
	if layerPrepTime == 0 {
		layerPrepTime += 2
	}
	totalTime := len(layers) * layerPrepTime
	return totalTime
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	numLayers := len(layers)
	for i := 0; i < numLayers; i++ {
		if layers[i] == "sauce" {
			sauce += 0.2
		}
		if layers[i] == "noodles" {
			noodles += 50
		}
	}
		return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fIngredients, oIngredients []string) {
	oIngredients[len(oIngredients) - 1] = fIngredients[len(fIngredients) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(recipe []float64, portions int) []float64 {
	scaled := make([]float64, len(recipe))
	copy(scaled, recipe)
	a := len(scaled)
	for i := 0; i < a; i++ {
		scaled[i] = scaled[i] / 2 * float64(portions)
	}
	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
