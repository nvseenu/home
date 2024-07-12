package hackerrank

import (
	"testing"
)

var timeInWordsTestData = []struct {
	hour   int32
	minute int32
	expRes string
}{
	{hour: 5, minute: 0, expRes: "five o' clock"},
	{hour: 5, minute: 1, expRes: "one minute past five"},
	{hour: 5, minute: 10, expRes: "ten minutes past five"},
	{hour: 5, minute: 15, expRes: "quarter past five"},
	{hour: 5, minute: 20, expRes: "twenty minutes past five"},
	{hour: 5, minute: 40, expRes: "twenty minutes to six"},
	{hour: 5, minute: 45, expRes: "quarter to six"},
	{hour: 5, minute: 47, expRes: "thirteen minutes to six"},
}

func TestTimeInWords(t *testing.T) {

	for _, td := range timeInWordsTestData {
		res := timeInWords(td.hour, td.minute)
		if res != td.expRes {
			t.Errorf("Expected [%v], but got [%v] for test data: %v", td.expRes, res, td)
		}
	}
}
