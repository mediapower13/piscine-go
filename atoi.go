package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	result := 0
	sign := 1
	i := 0

	// Check for sign
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		i = 1
	}

	// If string is only a sign
	if i >= len(s) {
		return 0
	}

	// Process digits
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	return result * sign
}
