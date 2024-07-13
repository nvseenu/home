package hackerrank

func ServiceLane(width []int32, entryExitPoints [][]int32) []int32 {

	res := []int32{}
	for _, entryExitPoint := range entryExitPoints {
		minWidth := int32(-1)
		for j := entryExitPoint[0]; j <= entryExitPoint[1]; j++ {
			if minWidth == -1 || width[j] < minWidth {
				minWidth = width[j]
			}
		}
		res = append(res, minWidth)
	}
	return res
}
