package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	skip := false

	for i := 0; i < len(str); i++ {
		// Skip spaces when counting characters
		if str[i] == ' ' {
			continue
		}

		// Skip one non-space character after every 5 characters
		if skip {
			skip = false
			continue
		}

		// Add non-space character to result
		result += string(str[i])
		count++

		// After 5 characters, add space and mark to skip next non-space
		if count == 5 {
			result += " "
			count = 0
			skip = true
		}
	}

	// Remove trailing space if exists and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result += "\n"

	return result
}
