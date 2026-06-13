func findMaxConsecutiveOnes(nums []int) int {
	maxOnes := 0
	streak := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			streak++
		} else {
			streak = 0
			continue
		}
		if streak > maxOnes {
			maxOnes = streak
		}
 	}
	return maxOnes
}
