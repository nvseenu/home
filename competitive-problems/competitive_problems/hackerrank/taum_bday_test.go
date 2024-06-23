package hackerrank

import "testing"

var taumTestData = []struct {
	b            int
	w            int
	bc           int
	wc           int
	z            int
	expectedCost int
}{
	{b: 3, w: 5, bc: 3, wc: 4, z: 1, expectedCost: 29},
	{b: 10, w: 10, bc: 1, wc: 1, z: 1, expectedCost: 20},
	{b: 5, w: 9, bc: 2, wc: 3, z: 4, expectedCost: 37},
	{b: 3, w: 6, bc: 9, wc: 1, z: 1, expectedCost: 12},
	{b: 7, w: 7, bc: 4, wc: 2, z: 1, expectedCost: 35},
	{b: 3, w: 3, bc: 1, wc: 9, z: 2, expectedCost: 12},
}

func TestMinCostToBuyGifts(t *testing.T) {
	for _, td := range taumTestData {
		actualCost := MinCostToBuyGifts(td.b, td.w, td.bc, td.wc, td.z)
		if actualCost != td.expectedCost {
			t.Errorf("Expected cost: %v, but got %v for test data: %v", td.expectedCost, actualCost, td)
		}
	}
}
