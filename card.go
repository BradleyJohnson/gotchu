package main

// Card contain info that's relevant to gameplay and aesthetics
type card struct {
	charValue  string
	position   int
	suit       string
	escapeCode string // for printing ansi colors
	pointValue int    // important for dragon/phoenix since they're worht 25/-25 respectively
}

func composeCard(suit string, position string) card {
	card := card{
		charValue:  position + " of " + suit,
		position:   cardRankMap[position],
		suit:       suit,
		escapeCode: escapeCodeMapping[suit],
		pointValue: pointValueMap[position],
	}
	return card
}

func composeWildCard(cardName string) card {
	wildCard := card{
		charValue:  cardName,
		position:   cardRankMap[cardName],
		suit:       "Wild",
		escapeCode: escapeCodeMapping["Wild"],
		pointValue: pointValueMap[cardName],
	}
	return wildCard
}

var escapeCodeMapping = map[string]string{
	"Jade":   "\u001b[32m%s\033[0m",
	"Star":   "\u001b[31m%s\033[0m",
	"Pagoda": "\u001b[33m%s\033[0m",
	"Sword":  "\u001b[36m%s\033[0m",
	"Wild":   "\u001b[34m%s\033[0m",
}

var cardRankMap = map[string]int{
	"Dog":      0,
	"Phoenix":  0,
	"Mah Jong": 1,
	"Two":      2,
	"Three":    3,
	"Four":     4,
	"Five":     5,
	"Six":      6,
	"Seven":    7,
	"Eight":    8,
	"Nine":     9,
	"Ten":      10,
	"Jack":     11,
	"Queen":    12,
	"King":     13,
	"Ace":      14,
	"Dragon":   20,
}

var pointValueMap = map[string]int{
	"Two":      0,
	"Three":    0,
	"Four":     0,
	"Five":     5,
	"Six":      0,
	"Seven":    0,
	"Eight":    0,
	"Nine":     0,
	"Ten":      10,
	"Jack":     0,
	"Queen":    0,
	"King":     10,
	"Ace":      0,
	"Dragon":   25,
	"Phoenix":  -25,
	"Dog":      0,
	"Mah Jong": 0,
}
