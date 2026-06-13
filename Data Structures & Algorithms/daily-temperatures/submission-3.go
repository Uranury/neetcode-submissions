type Temperature struct {
	Index int
	Temp int
}

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []Temperature{}

	for i, v := range temperatures {
		for len(stack) > 0 && v > stack[len(stack)-1].Temp {
			ind := stack[len(stack)-1].Index
			result[ind] = i - ind
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Temperature{Index: i, Temp: v})
	}

	return result
}
