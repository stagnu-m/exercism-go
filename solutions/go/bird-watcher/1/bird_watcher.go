package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, value := range birdsPerDay {
		sum += value
	}
	return sum
	// panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	for index, value := range birdsPerDay {
		if index > week*7 || index < (week-1)*7 {
			continue
		}
		sum += value
	}
	return sum
	// panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index]++
		}
	}
	return birdsPerDay
	// panic("Please implement the FixBirdCountLog() function")
}
