package piscine

func AtoiBase(s string, base string) int {
	// Check if base is valid
	if len(base) < 2 {
		return 0
	}

	// Check for + or - in base and duplicates
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		// Check for duplicates
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	// Convert string to integer
	result := 0
	baseLen := len(base)

	for _, char := range s {
		// Find position of char in base
		pos := -1
		for i, b := range base {
			if char == b {
				pos = i
				break
			}
		}

		// If char not in base, return 0
		if pos == -1 {
			return 0
		}

		result = result*baseLen + pos
	}

	return result
}
