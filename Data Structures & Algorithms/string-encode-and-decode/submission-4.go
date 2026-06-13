type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	encoded := ""
	for _, v := range strs {
		encoded += strconv.Itoa(len(v)) + "#" + v
	}
	return encoded
}

func (s *Solution )Decode(encoded string) []string {
	strs := make([]string, 0)
	r := []rune(encoded)

	i := 0
	j := 1

	for j < len(r) {
		length := 0
		if r[j] == '#' {
			length, _ = strconv.Atoi(string(r[i:j]))
			strs = append(strs, string(r[j+1:j+1+length]))
			i = j + length + 1
			j = i + 1
		} else {
			j++
		}
	}

	return strs
}
