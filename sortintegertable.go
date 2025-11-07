package piscine

// SortIntegerTable sorts the given slice of ints in ascending order in-place.
func SortIntegerTable(table []int) {
	n := len(table)
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
