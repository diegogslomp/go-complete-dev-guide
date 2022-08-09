package main

import (
	"fmt"
	"io"
	"os"
)

func Error(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		Error("Need filename as argument")
	}

	filename := os.Args[1]
	file, err := os.Open(filename)

	if err != nil {
		Error("Can not open file")
	}

	io.Copy(os.Stdout, file)
}
