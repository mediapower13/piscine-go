package piscine

import "github.com/01-edu/z01"

const BoardSize = 8

func EightQueens() {
	var queens [BoardSize]int
	digits := []rune{'1', '2', '3', '4', '5', '6', '7', '8'}

	printSolution := func() {
		for col := 0; col < BoardSize; col++ {
			z01.PrintRune(digits[queens[col]])
		}
		z01.PrintRune('\n')
	}

	isSafe := func(col, row int) bool {
		for c := 0; c < col; c++ {
			r := queens[c]
			if r == row {
				return false
			}
			dr := r - row
			if dr < 0 {
				dr = -dr
			}
			dc := c - col
			if dc < 0 {
				dc = -dc
			}
			if dr == dc {
				return false
			}
		}
		return true
	}

	var place func(col int)
	place = func(col int) {
		if col == BoardSize {
			printSolution()
			return
		}
		for row := 0; row < BoardSize; row++ {
			if isSafe(col, row) {
				queens[col] = row
				place(col + 1)
			}
		}
	}

	place(0)
}
