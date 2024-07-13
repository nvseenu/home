package hackerrank

import (
	"math"
)

func Workbook(chapters []int32, sumsPerPage int32) int32 {

	pageNum := int32(1)
	specialProblems := []int32{}
	for _, chapter := range chapters {
		pages := int32(math.Ceil(float64(chapter) / float64(sumsPerPage)))
		startPageNum := pageNum
		endPageNum := pageNum + pages - 1

		startSumNum := int32(1)
		for i := startPageNum; i <= endPageNum; i++ {
			endSumNum := startSumNum + sumsPerPage - 1
			if endSumNum > chapter {
				endSumNum = chapter
			}

			if i >= startSumNum && i <= endSumNum {
				specialProblems = append(specialProblems, int32(i))
			}
			startSumNum = endSumNum + 1
		}
		pageNum = endPageNum + 1

	}
	return int32(len(specialProblems))
}
