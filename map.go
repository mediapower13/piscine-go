package piscine

// Map applies the function f to each element of slice a and returns
// a new slice containing the boolean results in the same order.
func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, 0, len(a))
	for _, v := range a {
		res = append(res, f(v))
	}
	return res
}
