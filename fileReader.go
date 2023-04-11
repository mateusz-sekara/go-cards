package main

import (
	"fmt"
	"io"
	"os"
)

func fileSandbox() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments passed to command")
		os.Exit(1)
	}

	filename := os.Args[1]
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("Error while opening file:", err)
		os.Exit(1)
	}

	_, err = io.Copy(os.Stdout, f)
	if err != nil {
		fmt.Println("Error while copying file to sdout:", err)
		os.Exit(1)
	}
}
