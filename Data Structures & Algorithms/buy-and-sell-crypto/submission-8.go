func maxProfit(prices []int) int {
	maxprof := 0
	l, r := 0, 1
	for r < len(prices) && l < len(prices) {
		if prices[l] > prices[r] {
			l++
			if l < len(prices) - 1 {
				r = l + 1
				continue
			} else {
				break
			}
		}
		profit := prices[r] - prices[l]
		maxprof = max(maxprof, profit)

		r++
	}

	return maxprof
}
