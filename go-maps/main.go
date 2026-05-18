package main

import "fmt"

func print(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func main() {
	// colors := map[string]string{
	// 	"red":   "red",
	// 	"green": "green",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["red"] = "red"

	colors["10"] = "884884"

	delete(colors, "10")

	for key, value := range colors {
		fmt.Println(key, value)
	}

	fmt.Printf("%+v", colors)

	print(colors)
}
