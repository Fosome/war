package main

import (
	"fmt"
	"war/game"
)

func main() {
	standardDeck := game.NewStandardDeck()

	standardDeck.Shuffle()

	for i, card := range standardDeck.Cards() {
		fmt.Printf("Card %d: %s\n", i, card.String())
	}
}
