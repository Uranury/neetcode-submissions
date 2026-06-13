func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

	for i := range nums {
		complement := target - nums[i]
		if v, ok := m[complement]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	
	return nil
}
