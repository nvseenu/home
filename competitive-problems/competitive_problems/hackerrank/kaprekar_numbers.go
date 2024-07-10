package hackerrank

import (
	"strconv"
	"strings"
)

func FindKaprekrNumbers(startNum, endNum int32) string {
	res := []int{}
	for i := startNum; i <= endNum; i++ {
		if IsKaprekarNumber(int(i)) {
			res = append(res, int(i))
		}
	}

	if len(res) == 0 {
		return "INVALID RANGE"
	}

	numStrings := []string{}
	for _, c := range res {
		numStrings = append(numStrings, strconv.Itoa(c))
	}
	return strings.Join(numStrings, " ")

}

func IsKaprekarNumber(v int) bool {
	vstr := strconv.Itoa(v)
	vsquare := v * v
	vsquareStr := strconv.Itoa(vsquare)

	mid := len(vsquareStr) - len(vstr)
	firstPart := parseNumber(vsquareStr[:mid])
	secondPart := parseNumber(vsquareStr[mid:])
	return v == firstPart+secondPart
}

func parseNumber(s string) int {

	if v, err := strconv.Atoi(s); err != nil {
		return 0
	} else {
		return v
	}

}
