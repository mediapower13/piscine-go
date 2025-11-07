package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		a := os.Args[i]
		if a == "01" || a == "galaxy" || a == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
