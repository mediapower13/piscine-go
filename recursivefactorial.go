package piscine

// RecursiveFactorial returns the factorial of nb using recursion.
// For negative inputs or when the result would overflow int, it returns 0.
func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	// Prevent very deep recursion and integer overflow on typical platforms.
	// 20! fits in a signed 64-bit int, 21! does not. Return 0 for larger inputs.
	if nb > 20 {
		return 0
	}

	if nb == 0 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1)
}
