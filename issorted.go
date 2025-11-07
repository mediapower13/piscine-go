package piscine

// IsSorted returns true if slice a is sorted according to f.
// f returns >0 if first arg > second arg, 0 if equal, <0 otherwise.
func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n < 2 {
		return true
	}

	// Check non-decreasing (ascending): f(a[i], a[i+1]) <= 0 for all i
	asc := true
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			asc = false
			break
		}
	}
	if asc {
		return true
	}

	// Check non-increasing (descending): f(a[i], a[i+1]) >= 0 for all i
	desc := true
	for i := 0; i < n-1; i++ {
		if f(a[i], a[i+1]) < 0 {
			desc = false
			break
		}
	}
	return desc
}
