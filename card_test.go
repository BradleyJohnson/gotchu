package main

import "testing"

func TestComposeCard(t *testing.T) {
	card := composeCard("Pagoda", "Five")

	if card.charValue != "Five of Pagoda" {
		t.Errorf("\033[1;31m ERROR: Unexpected card name: card.charValue \033[0m")
	}

	if card.pointValue != 5 {
		t.Errorf("\033[1;31m ERROR: Expected point val of 5: got %v\033[0m", card.pointValue)
	}
}

func TestComposeWildCard(t *testing.T) {
	dragon := composeWildCard("Dragon")

	if dragon.pointValue != 25 {
		t.Errorf("\033[1;31m ERROR: Expected point val of 25 for Dragon: got %v\033[0m", dragon.pointValue)
	}
}
