package piscine

// RecursivePower returns nb raised to power using recursion.
// Negative powers return 0. Overflow is not handled per the exercise rules.
func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * RecursivePower(nb, power-1)
}
