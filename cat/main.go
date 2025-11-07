package main

import (
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		// read from stdin and copy to stdout
		_, _ = io.Copy(os.Stdout, os.Stdin)
		return
	}

	for _, name := range args {
		f, err := os.Open(name)
		if err != nil {
			// normalize error: avoid repeating "open <name>: " twice
			msg := err.Error()
			prefix := "open " + name + ": "
			if len(msg) >= len(prefix) && msg[:len(prefix)] == prefix {
				msg = msg[len(prefix):]
			}
			// write error to stderr in required format and exit immediately
			os.Stderr.WriteString("ERROR: open " + name + ": " + msg + "\n")
			os.Exit(1)
		}
		_, _ = io.Copy(os.Stdout, f)
		f.Close()
	}
}
