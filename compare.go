package piscine

// Compare compares two strings lexicographically like strings.Compare.
// It returns 0 if a==b, -1 if a<b, and 1 if a>b.
func Compare(a, b string) int {
	la, lb := len(a), len(b)
	min := la
	if lb < min {
		min = lb
	}
	for i := 0; i < min; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if la == lb {
		return 0
	} else if la < lb {
		return -1
	}
	return 1
}
