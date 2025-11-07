package piscine

// IsLower returns true if s contains only lowercase Latin letters (a-z) and is non-empty.
// Any other rune (space, punctuation, uppercase, digits, etc.) yields false.
func IsLower(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
