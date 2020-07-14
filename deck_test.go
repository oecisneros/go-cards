package main

import (
	"os"
	"testing"
)

func TestNewDesk(t *testing.T) {
	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected deck lenght of 16, but got %d", len(cards))
	}

	if cards[0].toString() != "Ace of Spades" {
		t.Errorf("Expected first card 'Ace of Spades', but got %s", cards[0])
	}

	if cards[len(cards)-1].toString() != "Four of Clubs" {
		t.Errorf("Expected last card 'Four of Clubs', but got %s", cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const filename string = "_decktesting"

	os.Remove(filename)

	cards := newDeck()
	cards.saveToFile(filename)

	loadedCards := newDeckFromFile(filename)

	if len(loadedCards) != len(cards) {
		t.Errorf("Expected deck lenght of %d, but got %d", len(cards), len(loadedCards))
	}

	os.Remove(filename)
}
