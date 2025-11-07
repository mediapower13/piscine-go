package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	// assign untyped rune '*' (42) directly to int without conversion
	ptr.x = '*'
	ptr.y = ptr.x / 2
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n/10 > 0 {
		printNbr(n / 10)
	}
	digit := n % 10
	ch := '0'
	i := 0
	for i < digit {
		ch++
		i++
	}
	z01.PrintRune(ch)
}

func main() {
	p := &point{}
	setPoint(p)
	printStr("x = ")
	printNbr(p.x)
	printStr(", y = ")
	printNbr(p.y)
	z01.PrintRune('\n')
}
