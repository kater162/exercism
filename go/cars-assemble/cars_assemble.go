package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var carsPerHour float64
	carsPerHour = float64(productionRate) * successRate / 100
	return carsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var carsPerMinute int
	carsPerMinute = int(float64(productionRate) / 60 * successRate / 100)
	return carsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const singleCar uint = 10000
	const tenCars uint = 95000
	var groups int = carsCount / 10
	var individual int = carsCount % 10
	var totalCost uint
	totalCost = uint(groups) * tenCars + uint(individual) * singleCar
	return totalCost
}
