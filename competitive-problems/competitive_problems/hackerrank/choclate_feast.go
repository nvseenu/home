package hackerrank

func ChocolateFeast(amount int32, barPrice int32, numOfWrappers int32) int32 {

	bars := amount / barPrice

	totalBars := bars
	wrappersAtHand := bars
	for wrappersAtHand >= numOfWrappers {
		bars = wrappersAtHand / numOfWrappers
		totalBars += bars
		wrappersAtHand -= bars * numOfWrappers
		wrappersAtHand += bars
	}

	return totalBars

}
