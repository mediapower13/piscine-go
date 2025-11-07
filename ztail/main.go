package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 || args[0] != "-c" {
		// nothing to do; exit with non-zero
		os.Exit(1)
	}

	var n int64
	if _, err := fmt.Sscanf(args[1], "%d", &n); err != nil || n < 0 {
		os.Exit(1)
	}
	files := args[2:]
	multiple := len(files) > 1
	hadError := false

	for i, name := range files {
		data, err := os.ReadFile(name)
		if err != nil {
			// print error and continue to next file
			fmt.Println(err.Error())
			hadError = true
			continue
		}
		if multiple {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", name)
		}
		size := int64(len(data))
		var start int64 = 0
		if n < size {
			start = size - n
		}
		os.Stdout.Write(data[start:])
	}

	if hadError {
		os.Exit(1)
	}
}
