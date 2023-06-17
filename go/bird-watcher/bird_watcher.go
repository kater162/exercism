package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	j := len(birdsPerDay)
	var birdsTotal int
	for i := 0; i < j; i++ {
		birdsTotal += birdsPerDay[i]
	}
	return birdsTotal
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
		starWeek := (week - 1) * 7
		endWeek := week * 7
		return TotalBirdCount(birdsPerDay[starWeek:endWeek])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	days := len(birdsPerDay)
	for i := 0; i < days; i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}
