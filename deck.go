package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck
// A deck is a slice of strings

type deck []string

/* Create and return a deck of playing cards */
func createDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Heards", "Clubs"}
	cardvalues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	/* Nested for loops */
	for _, suit := range cardSuits {
		for _, val := range cardvalues {
			cards = append(cards, val+" of "+suit)
		}
	}
	return cards
}

/* Print function */
func (d deck) printDeck() {

	for index, card := range d {
		fmt.Println(index, card)
	}
}

/* Deal a hand of size 'handSize' */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
