func maxArea(heights []int) int {
	max := 0
	j, k := 0, len(heights) - 1
	for j < k {
		if heights[j] > heights[k] {
			if heights[k] * (k-j) > max {
				max = heights[k] * (k-j)
			}
			k--
		} else {
			if heights[j] * (k-j) > max {
				max = heights[j] * (k-j)
			}
			j++
		}
	}
	return max
}

