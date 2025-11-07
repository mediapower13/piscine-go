package piscine

// LastRune returns the last rune of the string s.
// If s is empty, it returns rune(0).
func LastRune(s string) rune {
	var last rune
	for _, r := range s {
		last = r
	}
	return last
}
