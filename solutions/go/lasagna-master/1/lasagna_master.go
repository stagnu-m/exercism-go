package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		return len(layers) * 2
	}
	return len(layers) * averagePreparationTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int = 0
	var sauce float64 = 0

	for _, value := range layers {
		if value == "noodles" {
			noodles += 50
		}
		if value == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	for index, value := range myList {
		if value == "?" {
			myList[index] = friendsList[len(friendsList)-1]
		}
	}
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for index, value := range quantities {
		newQuantities[index] = value * float64(numberOfPortions) / 2
	}
	return newQuantities
}
