package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Slept for 2 seconds")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(4 * time.Second)
		fmt.Println("Slept for 4 seconds")
	}()

	fmt.Println("Waiting here!!!")
	wg.Wait()

	fmt.Println("Start small. Ship something.")
}
