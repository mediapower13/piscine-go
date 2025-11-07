package main

import (
	"os"

	"github.com/01-edu/z01"
)

// parsePositiveInt parses a string as a positive integer (no signs).
// Returns (value, true) on success, (0, false) on failure (non-digit or empty).
func parsePositiveInt(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	n := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}
	return n, true
}

func main() {
	if len(os.Args) <= 1 {
		return
	}

	args := os.Args[1:]
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	if len(args) == 0 {
		return
	}

	for _, a := range args {
		n, ok := parsePositiveInt(a)
		if !ok || n < 1 || n > 26 {
			z01.PrintRune(' ')
			continue
		}
		if upper {
			z01.PrintRune(rune('A' + (n - 1)))
		} else {
			z01.PrintRune(rune('a' + (n - 1)))
		}
	}
	z01.PrintRune('\n')
}
