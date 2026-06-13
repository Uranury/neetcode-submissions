func characterReplacement(s string, k int) int {
	longest := 0 
	l, r := 0, 0
	freq := make(map[byte]int)
	maxFreq := 0

	for r < len(s) {
		freq[s[r]]++
		if freq[s[r]] > maxFreq {
			maxFreq = freq[s[r]]
		}

		if (r - l + 1) - maxFreq > k {
			freq[s[l]]--
			l++
		}

		if r - l + 1 > longest {
			longest = r - l + 1
		}
		r++
	}
	

	return longest
}
