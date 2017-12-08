package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

/* Convert to string to save to file */
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}

/* Save byte slice of string to file */
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

/* Load deck of cards from file */
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

/* Shuffle cards */
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
