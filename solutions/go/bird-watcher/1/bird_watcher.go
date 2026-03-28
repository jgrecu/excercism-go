package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
    for _, b := range birdsPerDay {
        total += b
    }

    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	index := 7 * week
    return TotalBirdCount(birdsPerDay[index-7:index])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, n := range birdsPerDay {
        if i %2 == 0 {
        	birdsPerDay[i] = n + 1
        }
    }
    return birdsPerDay
}
