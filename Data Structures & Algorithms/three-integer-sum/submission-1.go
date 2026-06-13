func threeSum(nums []int) [][]int {
    results := [][]int{}
    sort.Ints(nums[:])

    for i := 0; i < len(nums); i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, len(nums) - 1
        for j < k {
            if -nums[i] == nums[j] + nums[k] {
                results = append(results, []int{nums[i], nums[j], nums[k]})
                j++
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
                for j < k && k+1 != len(nums) && nums[k] == nums[k+1] {
                    k--
                }
            } else if -nums[i] > nums[j] + nums[k] {
                j++
            } else if -nums[i] < nums[j] + nums[k] {
                k--
            }
        }
    }
    
    return results
}
