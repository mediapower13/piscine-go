package piscine

// CountIf returns the number of elements in tab for which f returns true.
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, s := range tab {
		if f(s) {
			count++
		}
	}
	return count
}
