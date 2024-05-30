package competitive_problems

import "fmt"

// Given the arrival and departure times of all trains that reach a railway station,
// the task is to find the minimum number of platforms required for the railway station so that no train waits.
// We are given two arrays that represent the arrival and departure times of trains that stop.

// Examples:

// Input: arr[] = {9:00, 9:40, 9:50, 11:00, 15:00, 18:00}, dep[] = {9:10, 12:00, 11:20, 11:30, 19:00, 20:00}
// Output: 3
// Explanation: There are at-most three trains at a time (time between 9:40 to 12:00)

// Input: arr[] = {9:00, 9:40}, dep[] = {9:10, 12:00}
// Output: 1
// Explanation: Only one platform is needed.

func FindingRequiredPlatforms(arrivalTimes []int, departureTimes []int) int {

	count := 0
	platformsRequired := 0
	ai := 0
	di := 0

	for ai < len(arrivalTimes) && di < len(departureTimes) {
		fmt.Println("count:", count, "ai", ai, "di", di, "platformsRequired", platformsRequired)
		if arrivalTimes[ai] < departureTimes[di] {
			count++
			ai++
			if count > platformsRequired {
				platformsRequired = count
			}
		} else {
			count--
			di++
		}
	}
	return platformsRequired
}
