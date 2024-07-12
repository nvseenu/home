package hackerrank

// Refer problem statement at https://www.hackerrank.com/challenges/halloween-sale/problem?isFullScreen=true

func FindMaxGamesToBuy(firstGamePrice int32, discountAfterFirstGame int32, minimumGamePrice int32, buyAmount int32) int32 {
	numOfGames := int32(0)
	amount := buyAmount

	if amount < firstGamePrice {
		return 0
	}

	for gameCost := firstGamePrice; ; gameCost -= discountAfterFirstGame {
		if gameCost <= minimumGamePrice {
			gameCost = minimumGamePrice
		}

		amount -= gameCost
		if amount >= 0 {
			numOfGames++
		} else {
			return numOfGames
		}
	}

}
