package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    var a float64
    a=(productionRate*successRate)/100
    return a
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    var a int 
    a = CalculateWorkingCarsPerHour/60
    return int(a)
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
    if carsCost >10{
        var a int 
    	a= carsCount%10
    	UpdatedCarsCount:=carsCount-a
    }elif carsCost =10{
    
    }else{
    
    }
	panic("CalculateCost not implemented")
}
