func lengthOfLastWord(s string) int {
	length := 0
	i := len(s) - 1
	for !isAlpha(s[i]) {
		i--
	}

	for i >= 0 {
		if !isAlpha(s[i]) {
			break
		}
		length++
		i--
	}

	return length
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
