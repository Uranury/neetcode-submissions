func lengthOfLongestSubstring(s string) int {
	visited := make(map[byte]bool)
	i := 0
	longest := 0
	
	for r := 0; r < len(s); r++ {
		for visited[s[r]] {
			delete(visited, s[i])
			i++
		}
		visited[s[r]] = true
		if r - i + 1 > longest {
			longest = r - i + 1
		}
	}

	return longest
}
