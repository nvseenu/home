package hackerrank

// Please refer problem statement at https://www.hackerrank.com/challenges/minimum-distances/problem?isFullScreen=true

func MinimumDistances(a []int32) int32 {

	nums := make(map[int32][]int)

	for i, v := range a {
		if _, ok := nums[v]; !ok {
			nums[v] = []int{i, -1}
		} else {
			nums[v][1] = i
		}
	}

	min := len(a)

	for _, v := range nums {
		if v[1] > -1 {
			val := v[1] - v[0]
			if val < min {
				min = val
			}
		}
	}
	if min == len(a) {
		return -1
	} else {
		return int32(min)
	}
}
