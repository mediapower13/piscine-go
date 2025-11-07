package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func sortRunes(r []rune) {
	n := len(r)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if r[j] < r[min] {
				min = j
			}
		}
		if min != i {
			r[i], r[min] = r[min], r[i]
		}
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}

	insertVal := ""
	order := false
	target := ""

	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--help" || a == "-h" {
			printHelp()
			return
		}
		if a == "--order" || a == "-o" {
			order = true
			continue
		}
		// handle --insert=val
		if len(a) > 3 && a[:3] == "-i=" {
			insertVal = a[3:]
			continue
		}
		if len(a) > 9 && a[:9] == "--insert=" {
			insertVal = a[9:]
			continue
		}
		if a == "--insert" || a == "-i" {
			if i+1 < len(args) {
				insertVal = args[i+1]
				i++
			}
			continue
		}
		// first non-flag argument is target
		if target == "" {
			target = a
		} else {
			// concatenate additional non-flag args
			target += a
		}
	}

	if target == "" && insertVal == "" && !order {
		printHelp()
		return
	}

	result := target
	if insertVal != "" {
		result = target + insertVal
		if target == "" {
			result = insertVal
		}
	}

	if order {
		r := []rune(result)
		sortRunes(r)
		result = string(r)
	}

	// print result using fmt (allowed) and also z01 not required here
	fmt.Println(result)
	_ = z01.PrintRune
}
