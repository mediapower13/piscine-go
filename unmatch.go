package piscine

// Unmatch returns the element that does not have a correspondent pair in a.
// If all numbers have pairs, returns -1.
func Unmatch(a []int) int {
	counts := make(map[int]int)
	for _, v := range a {
		counts[v]++
	}
	for _, v := range a {
		if counts[v]%2 == 1 {
			return v
		}
	}
	return -1
}
