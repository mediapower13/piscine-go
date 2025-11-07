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
		// Take 5 characters
		end := i + 5
		if end > len(noSpaces) {
			end = len(noSpaces)
		}
		result += noSpaces[i:end]
		i = end

		// Skip one character
		i++

		// Add space if there are more characters to process
		if i < len(noSpaces) {
			result += " "
		}
	}

	result += "\n"
	return result
}
