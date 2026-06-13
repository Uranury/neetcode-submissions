func groupAnagrams(strs []string) [][]string {
	sublists := [][]string{}
	visited := make(map[string]bool)
	for i := 0; i < len(strs); i++ {
		if _, ok := visited[strs[i]]; ok {
			continue
		}
		arr := []string{}
		arr = append(arr, strs[i])
		for j := i+1; j < len(strs); j++ {
			if checkAnagram(strs[i], strs[j]) {
				arr = append(arr, strs[j])
				visited[strs[j]] = true
			}
		}
		visited[strs[i]] = true
		sublists = append(sublists, arr)
	}
	return sublists
}

func checkAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)

	for i := range s { 
		m1[s[i]]++ 
		m2[t[i]]++ 
	}

	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}