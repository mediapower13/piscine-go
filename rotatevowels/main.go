package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + ('a' - 'A')
	}
	return r
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - ('a' - 'A')
	}
	return r
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	// join args with single spaces
	s := ""
	for i, a := range args {
		if i != 0 {
			s += " "
		}
		s += a
	}

	// collect vowels (keep original rune/case)
	var vowels []rune
	for _, r := range s {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}
	if len(vowels) == 0 {
		// print original args separated by spaces
		for i, a := range args {
			for _, r := range a {
				z01.PrintRune(r)
			}
			if i != len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// reverse vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// build output by replacing vowels with reversed ones (use original rune case from reversed list)
	vi := 0
	for _, r := range s {
		if isVowel(r) {
			v := vowels[vi]
			z01.PrintRune(v)
			vi++
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
