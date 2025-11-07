package piscine

func ShoppingListSort(slice []string) []string {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if len(slice[j]) < len(slice[minIdx]) {
				minIdx = j
			}
		}
		if minIdx != i {
			slice[i], slice[minIdx] = slice[minIdx], slice[i]
		}
	}
	return slice
}
