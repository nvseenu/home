package hackerrank

// Refer problem statement at https://www.hackerrank.com/challenges/taum-and-bday/problem?isFullScreen=true
func MinCostToBuyGifts(b int, w int, bc int, wc int, z int) int {
	totalCost := 0

	if bc == wc {
		totalCost = (b + w) * bc
	} else if bc+z <= wc {
		totalCost = (b+w)*bc + w*z
	} else if wc+z <= bc {
		totalCost = (b+w)*wc + b*z
	} else {
		totalCost = (b * bc) + (w * wc)
	}

	return totalCost

}
