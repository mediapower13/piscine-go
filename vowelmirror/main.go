package main

import (
	"os"

	"github.com/01-edu/z01"
)

func mirrorVowel(r rune) rune {
	// mapping a<->u, e<->o, i->i; preserve case
	lower := r
	if lower >= 'A' && lower <= 'Z' {
		lower = lower + ('a' - 'A')
	}
	var mapped rune
	switch lower {
	case 'a':
		mapped = 'u'
	case 'e':
		mapped = 'o'
	case 'i':
		mapped = 'i'
	case 'o':
		mapped = 'e'
	case 'u':
		mapped = 'a'
	default:
		return r
	}
	// restore case
	if r >= 'A' && r <= 'Z' {
		if mapped >= 'a' && mapped <= 'z' {
			mapped = mapped - ('a' - 'A')
		}
	}
	return mapped
}

func containsVowel(s string) bool {
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
	}
	return false
}

func process(s string) string {
	// if no vowels, return original
	if !containsVowel(s) {
		return s
	}
	// otherwise replace each vowel with its mirror
	runes := []rune(s)
	for i, r := range runes {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			runes[i] = mirrorVowel(r)
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		z01.PrintRune('\n')
		return
	}

	for i, a := range args {
		out := process(a)
		for _, r := range out {
			z01.PrintRune(r)
		}
		if i != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
