package piscine

// IterativePower returns nb raised to the power using an iterative loop.
// Negative powers return 0. Overflow is not handled per the exercise rules.
func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	res := 1
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}
