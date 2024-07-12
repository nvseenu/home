package hackerrank

import (
	"testing"
)

var minimumDistancesTestData = []struct {
	input  []int32
	expRes int32
}{
	{input: []int32{3, 2, 1, 2, 3}, expRes: 2},
	{input: []int32{7, 1, 3, 4, 1, 7}, expRes: 3},
	{input: []int32{7, 1, 3, 4}, expRes: -1},
}

func TestMinimumDistances(t *testing.T) {

	for _, td := range minimumDistancesTestData {
		res := MinimumDistances(td.input)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
