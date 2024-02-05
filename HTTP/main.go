package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	//////
	// bs := make([]byte, 99999) // the second argument shows the available slots on the created byte slice
	// The reason that we initialize this byte slice with this giant amount of capacity inside it, it's because the read function is not
	// set up to automatically resize the slice if the slice is already full

	// Body is the value that implements Reader
	// In the below line the response body should take the byte slice, takes its html that is inside the body, and read that
	// data into the byte slice bs
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	//////

	// But we dont need to do this whole thing anytime we wanna get data from body, so this is the alternative for the above approach:
	io.Copy(os.Stdout, resp.Body)
}

// logWriter implements Writer interface because it has the write method with the same input and output as the Writer
// inferface in this link: https://pkg.go.dev/io#Writer
func (logWriter) Write(bs []byte) (int, error) {
	
}