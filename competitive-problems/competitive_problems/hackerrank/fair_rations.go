package hackerrank

import "strconv"

func FairRations(breads []int32) string {
	// Write your code here
	bs := []int{}
	for _, v := range breads {
		bs = append(bs, int(v))
	}
	return fairRations(bs)
}

func fairRations(breads []int) string {
	breadCount := 0
	for i := 0; i < len(breads); i++ {
		if isOdd(breads[i]) {
			breads[i] = breads[i] + 1
			breadCount++
			if i+1 < len(breads) {
				breads[i+1] = breads[i+1] + 1
				breadCount++
			} else {
				return "NO"
			}
		}
	}
	return strconv.Itoa(breadCount)

}

func fairRations1(breads []int) string {

	oddNums := 0
	adjacentOddNums := 0

	for i := 0; i < len(breads); i++ {
		v := breads[i]
		if isOdd(v) {
			oddNums++
			if i-1 >= 0 && isOdd(breads[i-1]) {
				adjacentOddNums++
			}
			if i+1 < len(breads) && isOdd(breads[i+1]) {
				if adjacentOddNums > 0 {
					adjacentOddNums--
				}
			}
		}
	}
	if !isOdd(oddNums) {
		return strconv.Itoa((oddNums - adjacentOddNums) * 2)
	} else {
		return "NO"
	}

}

func isOdd(n int) bool {
	return n%2 != 0
}
