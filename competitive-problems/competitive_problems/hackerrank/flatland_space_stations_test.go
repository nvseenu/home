package hackerrank

import (
	"testing"
)

var flatAndSpaceStationsTestData = []struct {
	spceStations []int
	totalCities  int
	expRes       int
}{
	{spceStations: []int{1}, totalCities: 3, expRes: 1},
	{spceStations: []int{0, 4}, totalCities: 5, expRes: 2},
	{spceStations: []int{0, 1, 2, 4, 3, 5}, totalCities: 6, expRes: 0},
}

func TestFlatAndSpaceStations(t *testing.T) {

	for _, td := range flatAndSpaceStationsTestData {
		res := FlatlAndSpaceStations(td.totalCities, td.spceStations)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
