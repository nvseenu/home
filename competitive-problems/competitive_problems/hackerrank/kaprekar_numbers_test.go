package hackerrank

import (
	"fmt"
	"testing"
)

var kaprekarTestData = []struct {
	input  []int32
	expRes string
}{
	{input: []int32{1, 100}, expRes: "1 9 45 55 99"},
}

func TestFindKaprekarNums(t *testing.T) {

	s := []int{1, 2, 3, 4}
	fmt.Println("s", s[:2], s[2:])

	for _, td := range kaprekarTestData {
		res := FindKaprekrNumbers(td.input[0], td.input[1])
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
