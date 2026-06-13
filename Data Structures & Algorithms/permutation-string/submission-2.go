func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	freq1 := make(map[byte]int)
	for i, _ := range s1 {
		freq1[s1[i]]++
	}

	freq2 := make(map[byte]int)
	l, r := 0, len(s1) - 1
	fillMap(s2[l:r], freq2)

	for r < len(s2) {
		freq2[s2[r]]++
		if equal(freq1, freq2) {
			return true
		}
		freq2[s2[l]]--
		if freq2[s2[l]] == 0 {
			delete(freq2, s2[l])
		}
		l++
		r++
	}

	return false
}

func fillMap(s string, m map[byte]int) {
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
}

func equal(a, b map[byte]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		bv, ok := b[k]
		if !ok || bv != v {
			return false
		}
	}

	return true
}