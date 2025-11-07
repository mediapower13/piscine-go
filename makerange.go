package piscine

// MakeRange returns a slice containing all integers from min (inclusive)
// to max (exclusive). If min >= max, it returns nil.
// This implementation must not use append.
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	sz := max - min
	res := make([]int, sz)
	for i := 0; i < sz; i++ {
		res[i] = min + i
	}
	return res
}
