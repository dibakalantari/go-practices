package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// get the filename from the terminal executed command
	filename := os.Args[1]
	
	// open the file
	file , err := os.Open(filename)

	if(err != nil) {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	// Print file content to terminal
	io.Copy(os.Stdout, file)
}