func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
    stack := []rune{}
	m := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}

	for _, v := range s {
		if v == '[' || v == '(' || v == '{' {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 || v != m[stack[len(stack)-1]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
