package piscine

// NRune returns the nth rune (1-indexed) of s.
// If n is out of range (<=0 or > len(s) in runes), it returns rune(0).
func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}
	count := 1
	for _, r := range s {
		if count == n {
			return r
		}
		count++
	}
	return 0
}
