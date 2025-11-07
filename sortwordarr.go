package piscine

// SortWordArr sorts the slice of strings in-place by ASCII ascending order.
func SortWordArr(a []string) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
