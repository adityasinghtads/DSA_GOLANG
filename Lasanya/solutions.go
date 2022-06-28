package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	RemainingOvenTime := OvenTime - actualMinutesInOven
	return RemainingOvenTime
	panic("RemainingOvenTime not implemented")
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	if numberOfLayers > 0 {
		PreparationTime := (numberOfLayers * 2)
		return PreparationTime
	} else {
		PreparationTime := 0
		return PreparationTime
	}
	panic("PreparationTime not implemented")
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	ElapsedTime := (numberOfLayers * 2) + (actualMinutesInOven)
	return ElapsedTime
	panic("ElapsedTime not implemented")
}
