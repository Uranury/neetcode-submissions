func minEatingSpeed(piles []int, h int) int {
	start, end := 1, 0
	rate := 0

	for _, v := range piles {
		if v > end {
			end = v
		}
	}

	for start <= end {
		middle := (start + end) / 2 
		timeTaken := 0
		for _, v := range piles {
			timeTaken += (v + middle - 1) / middle 
		}
		if timeTaken > h {
			start = middle + 1
		} else {
			rate = middle
			end = middle - 1
		}
	}

	return rate
}
