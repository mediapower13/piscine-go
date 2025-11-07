package piscine

func Abort(a, b, c, d, e int) int {
	// Put all values in a slice
	arr := []int{a, b, c, d, e}

	// Sort the array using bubble sort
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// Return the middle element (median)
	return arr[2]
}
