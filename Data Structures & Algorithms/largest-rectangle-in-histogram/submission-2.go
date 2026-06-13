func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left, right := i - 1, i + 1
		area := heights[i]

		for left >= 0 && heights[left] >= heights[i] {
			area += heights[i]
			left--
		} 
		for right < len(heights) && heights[right] >= heights[i] {
			area += heights[i]
			right++
		}

		maxArea = max(maxArea, area)
	}

	return maxArea
}
