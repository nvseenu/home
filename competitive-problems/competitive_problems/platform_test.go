package competitive_problems

import "testing"

var trainTimingsTestData = []struct {
	desc                      string
	arrivalTimes              []int
	departureTimes            []int
	expectedPlatformsRequired int
}{
	{
		desc:                      "No overlap in train schedules",
		arrivalTimes:              []int{900, 940, 1500, 1900},
		departureTimes:            []int{910, 1200, 1800, 2000},
		expectedPlatformsRequired: 1,
	},
	{
		desc:                      "Overlapping train schedules",
		arrivalTimes:              []int{900, 940, 950, 1100, 1500, 1800},
		departureTimes:            []int{910, 1200, 1120, 1130, 1900, 2000},
		expectedPlatformsRequired: 3,
	},
}

func TestFindingPlatformsRequired(t *testing.T) {
	for _, td := range trainTimingsTestData {
		platformsRequired := FindingRequiredPlatforms(td.arrivalTimes, td.departureTimes)
		if td.expectedPlatformsRequired != platformsRequired {
			t.Errorf("Case: %s => Expected %v, got %v", td.desc, td.expectedPlatformsRequired, platformsRequired)
		}
	}
}
