func countSeniors(details []string) int {
    num := 0
	for _, s := range details {
		age, _ := strconv.Atoi(s[11:13])
		if age > 60 {
			num++
		}
	}
	return num
}