package hackerrank

import (
	"sort"
)

// Please refer problem statement https://www.hackerrank.com/challenges/bigger-is-greater/problem?isFullScreen=true

func BiggerIsGreater(w string) string {
	// Write your code here
	chars := []rune(w)
	idx1, idx2 := findIndex(chars)
	if idx1 == -1 || idx2 == -1 {
		return "no answer"
	}
	swap(chars, idx1, idx2)

	subString := chars[idx1+1:]
	sort.Slice(subString, func(i, j int) bool {
		return subString[i] < subString[j]
	})

	return string(chars)
}

func swap(chars []rune, i, j int) {
	t := chars[j]
	chars[j] = chars[i]
	chars[i] = t
}

func findIndex(chars []rune) (int, int) {
	maxi := -1
	maxj := -1
	for i := len(chars) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if chars[j] < chars[i] {
				if j > maxj {
					maxi = i
					maxj = j
				}
			}
		}
	}
	return maxj, maxi
}
