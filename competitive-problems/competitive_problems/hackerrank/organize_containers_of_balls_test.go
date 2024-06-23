package hackerrank

import "testing"

var containerTestData = []struct {
	containers [][]int32
	expRes     string
}{
	{containers: [][]int32{{1, 1}, {1, 1}}, expRes: "Possible"},
	{containers: [][]int32{{0, 2}, {1, 1}}, expRes: "Impossible"},
	{containers: [][]int32{{1, 3, 1}, {2, 1, 2}, {3, 3, 3}}, expRes: "Impossible"},
	{containers: [][]int32{{0, 2, 1}, {1, 1, 1}, {2, 0, 0}}, expRes: "Possible"},
}

func TestOranizeContainers(t *testing.T) {
	for _, td := range containerTestData {
		res := OrganizeContainers(td.containers)
		if res != td.expRes {
			t.Errorf("Expected cost: %v, but got %v for test data: %v", td.expRes, res, td)
		}
	}
}
