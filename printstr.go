package piscine

import "github.com/01-edu/z01"

// PrintStr prints the characters of s one by one.
func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
