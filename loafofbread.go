package piscine

func LoafOfBread(str string) string {

	if len(str) < 5 {
		return "Invalid Output\n"
	}

	var result string
	i := 0
	n := len(str)

	for i < n {
		var currentWord string

		for len(currentWord) < 5 && i < n {
			char := str[i]
			if char != ' ' {
				currentWord += string(char)
			}
			i++
		}

		result += currentWord

		if i < n {
			i++ // Skip the next character

			if i < n {
				result += " "
			}
		}
	}

	// Add the final newline
	result += "\n"
	return result
}
