package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)

	var word string
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result[word]++
			word = ""
		} else {
			word += string(str[i])
		}
	}
	result[word]++

	return result
}
