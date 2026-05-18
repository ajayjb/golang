package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, value := range d {
		fmt.Println(i, value)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	return strings.Split(string(bs), ",")
}

func (d *deck) shuffle() {
	for i := range *d {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		newPosistion := r.Intn(len(*d))
		(*d)[i], (*d)[newPosistion] = (*d)[newPosistion], (*d)[i]
	}
}
