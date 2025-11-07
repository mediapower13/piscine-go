package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for i := 0; i < len(str); i++ {
		// Skip spaces when counting characters
		if str[i] == ' ' {
			continue
		}

		// Add non-space character to result
		result += string(str[i])
		count++

		// After 5 characters, add space and skip next character
		if count == 5 {
			result += " "
			count = 0
			// Skip the next character (including spaces until we find a non-space to skip)
			i++
			for i < len(str) && str[i] == ' ' {
				i++
			}
			// Now we're at a non-space character, and the loop will increment i again
			// So we don't need to do anything else here
		}
	}

	// Remove trailing space if exists and add newline
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	result += "\n"

	return result
}
