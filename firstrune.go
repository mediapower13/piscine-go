package piscine

// FirstRune returns the first rune of the string s.
// If s is empty, it returns rune(0).
func FirstRune(s string) rune {
	for _, r := range s {
		return r
	}
	return 0
}
