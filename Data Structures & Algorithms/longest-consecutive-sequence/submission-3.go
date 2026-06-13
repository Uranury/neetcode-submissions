func longestConsecutive(nums []int) int {
	var longest int
	contains := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		contains[nums[i]] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		length := 1
		if _, ok := contains[nums[i]-1]; ok {
			continue
		}
		for {
			if _, ok := contains[nums[i]+length]; ok {
				length++
			} else {
				break
			}
		}
		if length > longest {
			longest = length
		}
	}

	return longest
}
