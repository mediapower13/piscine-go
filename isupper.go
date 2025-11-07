package piscine

// IsUpper returns true if s contains only uppercase Latin letters (A-Z) and is non-empty.
// Any other rune (space, punctuation, lowercase, digits, etc.) yields false.
func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
