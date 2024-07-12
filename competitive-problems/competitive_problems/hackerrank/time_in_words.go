package hackerrank

var numStrings = map[int]string{
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
	11: "eleven",
	12: "twelve",
	13: "thirteen",
	14: "fourteen",

	15: "fifteen",
	16: "sixteen",
	17: "seventeen",
	19: "ninteen",
	20: "twenty",
	30: "thirty",
	40: "fourty",
	50: "fifty",
}

func hourString(hour int32) string {
	return numStrings[int(hour)]
}

func minuteString(minute int32) string {
	if minute == 0 {
		return ""
	}
	if minute == 1 {
		return numStrings[int(minute)] + " minute"
	}
	if minute < 15 {
		return numStrings[int(minute)] + " minutes"
	} else if minute == 15 {
		return "quarter"
	} else if minute == 30 {
		return "half"
	} else {
		tens := minute / 10
		mstr := numStrings[int(tens*10)]
		ones := minute - (tens * 10)

		if ones == 0 {
			return mstr + " minutes"
		} else {
			return mstr + " " + numStrings[int(ones)] + " minutes"
		}
	}
}

func timeInWords(h int32, m int32) string {

	hourStr := hourString(h)
	minuteStr := minuteString(m)
	if m == 0 {
		return hourStr + " o' clock"
	} else if m <= 30 {
		return minuteStr + " past " + hourStr
	} else {
		return minuteString(60-m) + " to " + hourString(h+1)
	}

}
