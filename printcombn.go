package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	first := true
	printCombN(digits, n, 0, &first)
	z01.PrintRune('\n')
}

func printCombN(digits []int, n int, pos int, first *bool) {
	if pos == n {
		if !*first {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		*first = false
		for _, digit := range digits {
			z01.PrintRune('0' + rune(digit))
		}
		return
	}

	start := 0
	if pos > 0 {
		start = digits[pos-1] + 1
	}

	for i := start; i <= 9-(n-pos-1); i++ {
		digits[pos] = i
		printCombN(digits, n, pos+1, first)
	}
}
