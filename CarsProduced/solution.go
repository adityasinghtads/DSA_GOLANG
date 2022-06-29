package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var a float64
	a = float64(productionRate)
	x := (a * successRate) / 100
	return x
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var a float64
	var min float64
	min = 60
	a = CalculateWorkingCarsPerHour(productionRate, successRate) / min
	return int(a)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	if carsCount > 10 {
		var a int
		a = carsCount % 10
		UpdatedCarsCount := carsCount - a
		CalculateCost := ((UpdatedCarsCount / 10) * 95000) + (a * 10000)
		return uint(CalculateCost)
	} else if carsCount == 10 {
		CalculateCost := 95000
		return uint(CalculateCost)
	} else {
		CalculateCost := carsCount * 10000
		return uint(CalculateCost)
	}
	panic("CalculateCost not implemented")
}
