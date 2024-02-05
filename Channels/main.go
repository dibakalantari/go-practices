package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://go.dev",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for l := range c { // this is same as for{} , we are basically saying each time you receive something from the channel go through this loop
		// go checkLink(l, c) // we receive the link from the checkLink method and since we want to repeat checking status
		// of this links we keep creating new go routines to check the status as soon as one check is done for that link
	// }

	// sleeping a routine so they won't get called over and over immediately and we won't flood the websites with our requests
	for l := range c {
		// function literal in Go(anynomous function)
		go func (link string)  {
			time.Sleep(5 * time.Second)
		    checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if(err != nil) {
		fmt.Println(link,"might be broken")
		c <- link
		return
	}

	fmt.Println(link,"is ok")
	c <- link
}