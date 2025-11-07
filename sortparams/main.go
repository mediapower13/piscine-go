package main

import (
	"os"

	"github.com/01-edu/z01"
)

func lessASCII(a, b string) bool {
	ai := []byte(a)
	bi := []byte(b)
	for i := 0; i < len(ai) && i < len(bi); i++ {
		if ai[i] < bi[i] {
			return true
		}
		if ai[i] > bi[i] {
			return false
		}
	}
	return len(ai) < len(bi)
}

func main() {
	args := os.Args[1:]
	n := len(args)
	if n == 0 {
		return
	}

	// Simple selection sort using ASCII order (compare bytes)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if lessASCII(args[j], args[min]) {
				min = j
			}
		}
		if min != i {
			args[i], args[min] = args[min], args[i]
		}
	}

	for _, a := range args {
		for _, r := range a {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
