package hackerrank

import (
	"testing"
)

var tripletsTestData = []struct {
	arr    []int32
	d      int32
	expRes int32
}{
	{arr: []int32{2, 2, 3, 4, 5}, d: 1, expRes: 3},
	{arr: []int32{1, 2, 4, 5, 7, 8, 10}, d: 3, expRes: 3},
	{arr: []int32{1, 6, 7, 7, 8, 10, 12, 13, 14, 19}, d: 3, expRes: 2},
}

func TestBeautifulTripets(t *testing.T) {

	for _, td := range tripletsTestData {
		res := BeautifulTriplets(td.d, td.arr)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
