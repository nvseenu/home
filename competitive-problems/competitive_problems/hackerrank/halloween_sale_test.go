package hackerrank

import (
	"testing"
)

var halloweenSaleTestData = []struct {
	firstGamePrice     int32
	secondGameDiscount int32
	minimumGamePrice   int32
	buyAmount          int32
	expRes             int32
}{
	{firstGamePrice: 20, secondGameDiscount: 3, minimumGamePrice: 6, buyAmount: 70, expRes: 5},
	{firstGamePrice: 20, secondGameDiscount: 3, minimumGamePrice: 6, buyAmount: 80, expRes: 6},
	{firstGamePrice: 20, secondGameDiscount: 3, minimumGamePrice: 6, buyAmount: 20, expRes: 1},
	{firstGamePrice: 20, secondGameDiscount: 3, minimumGamePrice: 6, buyAmount: 19, expRes: 0},
}

func TestHalloweenSale(t *testing.T) {

	for _, td := range halloweenSaleTestData {
		res := FindMaxGamesToBuy(td.firstGamePrice, td.secondGameDiscount, td.minimumGamePrice, td.buyAmount)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
