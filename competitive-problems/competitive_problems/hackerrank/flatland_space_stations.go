package hackerrank

import (
	"math"
	"sort"
)

func FlatlAndSpaceStations(totalCities int, stations []int) int {
	sort.Slice(stations, func(i, j int) bool {
		return stations[i] < stations[j]
	})
	distances := []int{}
	spaceIndex := 0
	for c := 0; c < totalCities; c++ {

		s := spaceStation(stations, spaceIndex)
		d := int(math.Abs(float64(c - s)))
		ns := spaceStation(stations, spaceIndex+1)
		nd := int(math.Abs(float64(c - ns)))

		nearDistance := 0
		if nd <= d {
			nearDistance = nd
			spaceIndex++
		} else {
			nearDistance = d
		}
		distances = append(distances, nearDistance)

	}

	maxDistance := 0
	for _, v := range distances {
		if v > maxDistance {
			maxDistance = v
		}
	}
	return maxDistance
}

func spaceStation(stations []int, index int) int {
	if index >= len(stations) {
		index = len(stations) - 1
	}
	return stations[index]
}
