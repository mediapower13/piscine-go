package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Cannot read file")
		return
	}
	defer f.Close()
	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Error reading file")
	}
}
