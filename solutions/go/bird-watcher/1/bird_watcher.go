package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := range len(birdsPerDay){
		sum += birdsPerDay[i]
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// sum := 0
	// left := (week-1) * 7 + 0
	// right := (week-1) * 7 + 6
	// for i := range len(birdsPerDay){
	// 	if i >= left && i <= right{
	// 		sum += birdsPerDay[i]
	// 	}
	// }
	// return sum
	start := (week - 1) * 7
	end := min(start+7, len(birdsPerDay))

	sum := 0
	for _, count := range birdsPerDay[start:end] {
		sum += count
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range(len(birdsPerDay)){
		if i%2 == 0{
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
