func maxProfit(prices []int) int {
	i, j := 0, 1
	profit := 0
	for i < len(prices) && j < len(prices) {
		if prices[i] < prices[j] {
			if prices[j] - prices[i] > profit {
				profit = prices[j] - prices[i]
			}
			j++
		} else {
			i++
			j = i + 1
		}
	}	

	return profit
}
