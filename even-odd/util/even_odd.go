package util

import "fmt"

func EvenOrOdd() {
	array := []int{}

	for i := range 11 {
		array = append(array, i)
	}

	fmt.Println(array)

	for i := range 11 {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("odd")
		}
	}

}
