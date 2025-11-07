package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	n := 0
	for range deck {
		n++
	}
	if n == 0 {
		return
	}
	players := 4
	per := n / players
	for p := 0; p < players; p++ {
		fmt.Printf("Player %d: ", p+1)
		for i := 0; i < per; i++ {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%d", deck[p*per+i])
		}
		fmt.Printf("\n")
	}
}
