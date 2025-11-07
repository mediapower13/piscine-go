package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0
	i := 0

	for i < len(str) {
		// Skip spaces
		if str[i] == ' ' {
			i++
			continue
		}

		// Add character to result
		result += string(str[i])
		count++
		i++

		// After 5 characters, add space and skip next non-space character
		if count == 5 {
			result += " "
			count = 0
			// Skip the next non-space character
			for i < len(str) && str[i] == ' ' {
				i++
			}
			if i < len(str) {
				i++ // Skip one non-space character
			}
		}
	}

	// Remove trailing space if exists and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result += "\n"

	return result
}
