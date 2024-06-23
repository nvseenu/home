package hackerrank

import "testing"

var encryptionTestData = []struct {
	input  string
	expRes string
}{
	{input: "feedthedog", expRes: "fto ehg ee dd"},
	{input: "feed the doger", expRes: "fto ehg eee ddr"},

	{input: "haveaniceday", expRes: "hae and via ecy"},
	{input: "if man was meant to stay on the ground god would have given us roots", expRes: "imtgdvs fearwer mayoogo anouuio ntnnlvt wttddes aohghn sseoau"},
}

func TestEncryption(t *testing.T) {

	for _, td := range encryptionTestData {
		res := Encryption(td.input)
		if res != td.expRes {
			t.Errorf("Expected >%v<, but got >%v<for test data: %v", td.expRes, res, td)
		}
	}
}
