func topKFrequent(nums []int, k int) []int {
	result := []int{}
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	freqArr := make(map[int][]int)
	for k, v := range freq {
		freqArr[v] = append(freqArr[v], k)
	}

	for i := len(nums); i > 0; i-- {
		if len(result) == k {
			break
		}
		if arr, ok := freqArr[i]; ok {
			for _, v := range arr {
				if len(result) == k {
					break
				}
				result = append(result, v)
			}
		}
	}
	return result
}
