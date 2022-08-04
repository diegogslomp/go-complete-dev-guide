package main

import (
	"fmt"
	"io"
	"os"
)

func printError(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		printError("Need filename as argument")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		printError("Can not open file")
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
