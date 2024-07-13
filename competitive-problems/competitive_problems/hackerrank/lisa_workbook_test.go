package hackerrank

import (
	"testing"
)

var workbookTestData = []struct {
	chapters []int32
	problems int32
	expRes   int32
}{
	{chapters: []int32{4, 2}, problems: 3, expRes: 1},
	{chapters: []int32{4, 2, 6, 1, 10}, problems: 3, expRes: 4},
}

func TestWorkbook(t *testing.T) {

	for _, td := range workbookTestData {
		res := Workbook(td.chapters, td.problems)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
