func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	prefix := make([]int, len(nums))
	prefix[0] = 1 
	res := 1

	for i := 1; i < len(nums); i++ {
		res *= nums[i-1]
		prefix[i] = res
	}

	suffix := make([]int, len(nums))
	suffix[len(nums)-1] = 1

	for i := len(nums) - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}
