package hackerrank

// Pleae refer problem statement at https://www.hackerrank.com/challenges/beautiful-triplets/problem?isFullScreen=true

func BeautifulTriplets(d int32, arr []int32) int32 {
	res := [][]int{}

	i := 0
	j := i + 1
	k := j + 1
	for i < len(arr)-2 {
		if k >= len(arr) {
			i++
			j = i + 1
			k = j + 1
			continue
		}

		d1 := arr[j] - arr[i]
		d2 := arr[k] - arr[j]

		if d1 > d {
			i++
			j = i + 1
			k = j + 1

		} else if d1 == d {
			if d2 > d {
				j++
				k = j + 1
			} else if d2 == d {
				res = append(res, []int{i, j, k})
				i++
				j = i + 1
				k = j + 1
			} else {
				k++
			}

		} else {
			j++
			k++
		}
	}
	return int32(len(res))
}
