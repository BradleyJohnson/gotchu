package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suits = [4]string{"Jade", "Pagoda", "Sword", "Star"}
var positions = [13]string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}
var wild_cards = [4]string{"Mah Jong", "Dog", "Phoenix", "Dragon"}

type deck []card

func newDeck() deck {
	deck := deck{}

	for _, suit := range suits {
		for _, position := range positions {
			deck = append(deck, composeCard(suit, position))
		}
	}

	for _, card := range wild_cards {
		deck = append(deck, composeWildCard(card))
	}

	return deck
}

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card.charValue)
	}
}

func (d deck) shuffle() deck {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
	return d
}
