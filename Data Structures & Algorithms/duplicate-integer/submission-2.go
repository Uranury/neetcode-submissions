func hasDuplicate(nums []int) bool {
	m := make(map[int]bool)
	result := false

    for _, v := range nums {
		_, ok := m[v]
		if ok == false {
			m[v] = true
			continue
		}
		result = true
	}
	return result
}
