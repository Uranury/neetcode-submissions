func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left, right := -1, -1
		area := heights[i]
		if i > 0 {
			left = i - 1
		} 
		if i + 1 < len(heights) {
			right = i + 1
		}

		for left >= 0 && heights[left] >= heights[i] {
			area += heights[i]
			left--
		} 
		for right < len(heights) && right > 0 && heights[right] >= heights[i] {
			area += heights[i]
			right++
		}

		maxArea = max(maxArea, area)
	}

	return maxArea
}
