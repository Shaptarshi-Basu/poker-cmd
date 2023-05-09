package main

import (
	"fmt"
	"os"
	"poker/operations"
)

func main() {
	args := os.Args[1:] // Get all command-line arguments except for the program name

	if len(args) < 2 {
		fmt.Println("Please provide two string arguments as the two hands of the poker.")
		return
	}

	fh := args[0]
	sh := args[1]
	operations.HandValidator(fh, sh)
}
