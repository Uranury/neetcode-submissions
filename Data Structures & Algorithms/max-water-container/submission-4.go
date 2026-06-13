func maxArea(heights []int) int {
	max := 0
	for j := 0; j < len(heights); j++ {
		for k := len(heights) - 1; k > j; k-- {
			smaller := 0
			if heights[j] > heights[k] {
				smaller = heights[k]
			} else {
				smaller = heights[j]
			}
			if smaller * (k-j) > max {
				max = smaller * (k-j)
			}
		}
	} 
	return max
}
