package hackerrank

import (
	"testing"
)

var fairRationsTestData = []struct {
	input  []int32
	expRes string
}{
	{input: []int32{1, 2}, expRes: "NO"},
	{input: []int32{4, 5, 6, 7}, expRes: "4"},
	{input: []int32{4, 4, 7, 7}, expRes: "2"},
}

func TestFairRations(t *testing.T) {

	for _, td := range fairRationsTestData {
		res := FairRations(td.input)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
