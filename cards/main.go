package main

import (
	"fmt"
)

func main() {
	cards := newDeck()

	cards.print()

	cards.shuffle()

	fmt.Print("\n")

	cards.print()

	// num := "TEST"
	// p := &num

	// fmt.Println(num)
	// fmt.Println(p)
	// fmt.Println(*p)

	// fmt.Println(readDeckFromFile("my_caurds.txt"))

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for i := range 10 {
	// 	fmt.Println(i)
	// }

	// for index, value := range cards {
	// 	fmt.Println(index, value)
	// }

	// txt := "New Brave World"

	// for _, char := range txt {
	// 	fmt.Println(string(char), "beast")
	// }

	// for i := range txt {
	// 	fmt.Println(i)
	// }
}

func newCard() string {
	return "Five of Diamonds"
}
