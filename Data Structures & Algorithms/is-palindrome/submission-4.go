func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s) - 1
	for i < j {
		if !isAlphaNumeric(rune(s[i])) {
			i++
			continue
		}
		if !isAlphaNumeric(rune(s[j])) {
			j--
			continue
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}