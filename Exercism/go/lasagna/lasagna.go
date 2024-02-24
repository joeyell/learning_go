package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	timeRemainingInOven := OvenTime - actualMinutesInOven
	return timeRemainingInOven
	//panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	prepTime := 2 * numberOfLayers
	return prepTime
	//panic("PreparationTime not implemented")
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	totalTime := actualMinutesInOven + PreparationTime(numberOfLayers)
	return totalTime
	//panic("ElapsedTime not implemented")
}
