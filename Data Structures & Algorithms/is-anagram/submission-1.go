func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := range s {
		m1[s[i]]++
		m2[t[i]]++
	}
	
	for v, k := range m1 {
		if m2[v] != k {
			return false
		}
	}
	return true
}
