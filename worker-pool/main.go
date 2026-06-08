package main

import (
	"fmt"
	"sync"
	"time"
)

func sendSignal() <-chan int {
	ch := make(chan int, 100)

	go func() {
		defer close(ch)
		for i := range 100 {
			// time.Sleep(0.5 * time.Second)
			ch <- i
		}
	}()

	return ch
}

func main() {
	workers := 4
	sem := make(chan struct{}, workers)
	var wg sync.WaitGroup

	ch := sendSignal()

	for data := range ch {
		fmt.Println("Processing Data Start", data)
		sem <- struct{}{}

		//   wg.Go(func(){
		//       defer func(){<- sem}()
		//       fmt.Println("Processing Data Start", data)
		//       time.Sleep(5 * time.Second)
		//       fmt.Println("Processing Data End", data)
		//   })

		// or

		wg.Add(1)
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			fmt.Println("Processing Data Start", data)
			time.Sleep(5 * time.Second)
			fmt.Println("Processing Data End", data)
		}()
	}

	wg.Wait()
}
