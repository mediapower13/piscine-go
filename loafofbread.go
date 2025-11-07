package piscine

func LoafOfBread(str string) string {
	// Remove all spaces first to simplify
	noSpaces := ""
	for _, char := range str {
		if char != ' ' {
			noSpaces += string(char)
		}
	}

	if len(noSpaces) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	for i < len(noSpaces) {
		// Take up to 5 characters
		count := 0
		for count < 5 && i < len(noSpaces) {
			result += string(noSpaces[i])
			i++
			count++
		}

		// After taking 5, skip one character based on remaining
		if count == 5 && i < len(noSpaces) {
			remaining := len(noSpaces) - i
			if remaining >= 6 || remaining == 2 {
				i++ // Skip this character
			}
		}

		// Add space if there are more characters to process
		if i < len(noSpaces) {
			result += " "
		}
	}

	result += "\n"
	return result
}
