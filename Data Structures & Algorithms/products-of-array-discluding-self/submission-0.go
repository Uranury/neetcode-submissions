func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res := 1
		for j := 0; j < i; j++ {
			res *= nums[j]
		}
		for k := i+1; k < len(nums); k++ {
			res *= nums[k]
		}
		result[i] = res 
	}
	return result
}
