package piscine

// IsNumeric returns true if s contains only numerical ASCII characters (0-9) and is non-empty.
func IsNumeric(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
