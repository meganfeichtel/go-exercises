package main

import (
	"os"
	"testing"
)

var cardsInDeck = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != cardsInDeck {
		t.Errorf("Epected deck length of %v, but got %v", cardsInDeck, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Epected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Epected first card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != cardsInDeck {
		t.Errorf("Epected deck length of %v, but got %v", cardsInDeck, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Epected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Epected first card of Ace of Spades, but got %v", d[len(d)-1])
	}

	os.Remove("_decktesting")
}
