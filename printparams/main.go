package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Print all provided parameters, one per line
	args := os.Args[1:]
	for _, a := range args {
		for _, r := range a {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
