package birdwatcher
import "fmt"

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var sum = 0
    for i := 0; i < len(birdsPerDay); i++ {
        sum += birdsPerDay[i]
    }
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var sum = 0
    var end = 7 * (week)
    fmt.Println(end)
    var start = 0
    if(week == 1) { 
        start = 0
        end = 7
    } else {
        start = end - 7
    }
    fmt.Println(start)
    fmt.Println(end)
    var slice []int = birdsPerDay[start:end]
    for i := 0; i < len(slice) ; i++{
        sum += slice[i]
    }
    fmt.Println(sum)
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0 ; i < len(birdsPerDay) ; i++ {
        if(i % 2 == 0) {
            birdsPerDay[i] += 1
        }
    }
    return birdsPerDay
}
