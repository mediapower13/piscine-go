package piscine

import "github.com/01-edu/z01"

// PrintNbrBase prints nbr in the provided base. If the base is invalid, prints "NV".
// Validity: base must have at least 2 unique characters and must not contain '+' or '-'.
func PrintNbrBase(nbr int, base string) {
	// validate base
	br := []rune(base)
	if len(br) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	seen := make(map[rune]bool)
	for _, r := range br {
		if r == '+' || r == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		if seen[r] {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		seen[r] = true
	}

	// handle zero
	if nbr == 0 {
		z01.PrintRune(br[0])
		return
	}

	// build digits (handle negative without negating the whole number)
	var digits []rune
	baseLen := len(br)
	n := nbr
	for n != 0 {
		rem := n % baseLen
		if rem < 0 {
			rem = -rem
		}
		digits = append(digits, br[rem])
		n = n / baseLen
	}

	// print sign if negative
	if nbr < 0 {
		z01.PrintRune('-')
	}

	// print digits in reverse
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}
