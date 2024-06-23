package hackerrank

import (
	"math"
	"strings"
)

// Refer problem statement at https://www.hackerrank.com/challenges/encryption/problem?isFullScreen=true

func Encryption(s string) string {

	//Remove spaces
	val := strings.ReplaceAll(s, " ", "")
	res := math.Sqrt(float64(len(val)))
	colLen := int(math.Ceil(res))

	target := make([]byte, len(val)+colLen-1)
	ti := 0
	for ri := 0; ri < colLen; ri++ {
		for ci := 0; ci < colLen; ci++ {
			si := ri + (ci * colLen)
			if si < len(val) {
				target[ti] = val[si]
				ti++
			}
		}
		if ri+1 < colLen {
			target[ti] = 32
			ti++
		}
	}

	return string(target)

}
