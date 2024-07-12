package hackerrank

import "testing"

var choclateFeastTestData = []struct {
	amount   int32
	barPrice int32
	wrappers int32
	expRes   int32
}{
	{amount: 15, barPrice: 3, wrappers: 2, expRes: 9},
	{amount: 3, barPrice: 3, wrappers: 2, expRes: 1},
	{amount: 15, barPrice: 3, wrappers: 3, expRes: 7},
	{amount: 10, barPrice: 2, wrappers: 5, expRes: 6},
	{amount: 12, barPrice: 4, wrappers: 4, expRes: 3},
}

func TestChoclateFeast(t *testing.T) {

	for _, td := range choclateFeastTestData {
		res := ChocolateFeast(td.amount, td.barPrice, td.wrappers)
		if res != td.expRes {
			t.Errorf("Expected [%v], but got [%v] for test data: %v", td.expRes, res, td)
		}
	}
}
