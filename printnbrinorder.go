package piscine

import "github.com/01-edu/z01"

// PrintNbrInOrder prints the digits of n in ascending order.
// n is expected to be non-negative. Handles 0 and all int values without converting to int64.
func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var count [10]int
	for n > 0 {
		d := n % 10
		count[d]++
		n = n / 10
	}
	for d := 0; d <= 9; d++ {
		for i := 0; i < count[d]; i++ {
			z01.PrintRune('0' + rune(d))
		}
	}
}
