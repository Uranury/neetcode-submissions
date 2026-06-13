func maxArea(heights []int) int {
	l, r := 0, len(heights) - 1 
	maxArea := 0

	for l < r {
		area := min(heights[l], heights[r]) * (r - l)

		maxArea = max(maxArea, area)

		if heights[l] > heights[r] {
			r--
		} else {
			l++
		}
	}

	return maxArea
}
