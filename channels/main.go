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
		"http://golang.org",
		"http://amazon.com",
	}

	d := make(chan string, 2)

	d <- "Its time to leave planet"

	fmt.Println((<-d))

	// Above code only works with buffered channel if use with unbuffered channel we will get fatal error "fatal error: all goroutines are asleep - deadlock!"
	// If we use buffered channel we must use the go routines

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Yeah it might be down!")
		c <- "Yeah it might be down!"
		return
	}
	fmt.Println(link)
	c <- link
}
