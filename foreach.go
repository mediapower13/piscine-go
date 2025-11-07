package piscine

// ForEach applies the function f to every element of slice a.
func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
