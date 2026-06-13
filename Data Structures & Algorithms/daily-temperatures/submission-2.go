type Temperature struct {
	Index int
	Temp int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []Temperature{}

	for i, v := range temperatures {
		for len(stack) != 0 && stack[len(stack)-1].Temp < v {
			stackInd := stack[len(stack)-1].Index
			stack = stack[:len(stack)-1]
			result[stackInd] = (i - stackInd)
		}
		stack = append(stack, Temperature{Index: i, Temp: v})
	} 

	return result
}
