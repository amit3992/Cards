package main

import (
	"fmt"
)

func main() {
	cards := createDeck()
	// cards.printDeck()

	hand, cards := deal(cards, 5)

	fmt.Println(hand)

	fmt.Println(cards.toString())

}
