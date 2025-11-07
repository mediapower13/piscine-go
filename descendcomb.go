package piscine

import "github.com/01-edu/z01"

// DescendComb prints all combinations of two different two-digit numbers in descending order.
func DescendComb() {
	first := true
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false

			// Print first number (a)
			z01.PrintRune(rune('0' + a/10))
			z01.PrintRune(rune('0' + a%10))
			z01.PrintRune(' ')

			// Print second number (b)
			z01.PrintRune(rune('0' + b/10))
			z01.PrintRune(rune('0' + b%10))
		}
	}
	// Note: do not print a trailing newline so the output matches the test harness exactly
}
