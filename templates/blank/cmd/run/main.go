package main

import (
	"fmt"
	"os"

	"function"
)

func main() {
	if err := function.Run(); err != nil {
		fmt.Fprintf(os.Stderr, err)
		os.Exit(1)
	}
}
