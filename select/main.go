package main

import "fmt"

func main() {
	fmt.Println("Start small. Ship something.")

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 4
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case c1, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println(c1)
		case c2, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println(c2)
		default:
			fmt.Println("No channel found!")
		}
	}
}
