package hackerrank

import (
	"reflect"
	"testing"
)

var serviceLaneTestData = []struct {
	width  []int32
	cases  [][]int32
	expRes []int32
}{

	{width: []int32{2, 3, 1, 2, 3, 2, 3, 3}, cases: [][]int32{{0, 3}, {4, 6}, {6, 7}, {3, 5}, {0, 7}}, expRes: []int32{1, 2, 3, 2, 1}},
	{width: []int32{2, 3, 2, 1}, cases: [][]int32{{0, 1}, {1, 3}}, expRes: []int32{2, 1}},
}

func TestServiceLanes(t *testing.T) {

	for _, td := range serviceLaneTestData {
		res := ServiceLane(td.width, td.cases)
		if !reflect.DeepEqual(td.expRes, res) {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
