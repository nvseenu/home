package hackerrank

import "testing"

var biggerIsGreterTestData = []struct {
	input  string
	expRes string
}{
	{input: "abcd", expRes: "abdc"},
	{input: "dcba", expRes: "no answer"},
	{input: "bb", expRes: "no answer"},
	{input: "dkhc", expRes: "hcdk"},
	{input: "ehdegnmorgafrjxvksc", expRes: "ehdegnmorgafrjxvsck"},
}

func TestBiggerIsGreater(t *testing.T) {

	for _, td := range biggerIsGreterTestData {
		res := BiggerIsGreater(td.input)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
