package hackerrank

// Refer problem statement at https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem?isFullScreen=true

func OrganizeContainers(containers [][]int32) string {
	containerCapacities := make([]int32, len(containers))
	itemCounts := make([]int32, len(containers[0]))
	const CONTAINER_USED = -1

	for i := 0; i < len(containers); i++ {
		cap := int32(0)
		for j := 0; j < len(containers[i]); j++ {
			cap += containers[i][j]
			itemCounts[j] += containers[i][j]

		}
		containerCapacities[i] = cap
	}

	for _, itemCount := range itemCounts {
		for ci, cap := range containerCapacities {
			if itemCount == cap {
				// -1 indicates that particular container has been used,
				// and should not be referred for future comparison
				containerCapacities[ci] = CONTAINER_USED
			}
		}
	}

	count := 0
	for _, cap := range containerCapacities {
		if cap == CONTAINER_USED {
			count++
		}
	}

	if count == len(containerCapacities) {
		return "Possible"
	} else {
		return "Impossible"
	}
}
