package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Spades of Ace" {
		t.Errorf("Expected first card is Spades of Ace, but got %v", d[0])
	}

	if d[len(d)-1] != "Clubs of Three" {
		t.Errorf("Expected first card is Clubs of Three, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_deck_testing.txt")

	d := newDeck()

	d.saveToFile("_deck_testing.txt")

	loadedDeck := readDeckFromFile("_deck_testing.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	os.Remove("_deck_testing.txt")
}
