package piscine

// Swap swaps the values pointed to by a and b.
func Swap(a *int, b *int) {
	if a == nil || b == nil {
		return
	}
	tmp := *a
	*a = *b
	*b = tmp
}
