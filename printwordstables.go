package piscine

import "github.com/01-edu/z01"

// PrintWordsTables prints each string in the slice on its own line.
func PrintWordsTables(a []string) {
	for _, s := range a {
		for _, r := range s {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
