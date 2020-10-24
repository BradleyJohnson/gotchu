package main

import (
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 56 {
		t.Errorf("\033[1;31m ERROR: Expected newDeck() to return 56 cards, got: %d \033[0m", len(deck))
	}

	missingDragon := true
	for _, card := range deck {
		if card.charValue == "Dragon" {
			missingDragon = false
		}
	}
	if missingDragon == true {
		t.Errorf("\033[1;31m ERROR: The Dragon is missing from the deck \033[0m")
	}
}

func TestShuffle(t *testing.T) {
	deck := newDeck().shuffle()
	deckTwo := newDeck().shuffle()
	deckEquality := true

	for i := range deck {
		if deck[i] != deckTwo[i] {
			deckEquality = false
			break
		}
	}

	if deckEquality == true {
		t.Errorf("\033[1;31m ERROR: Two decks shuffled identically \033[0m")
	}
}
