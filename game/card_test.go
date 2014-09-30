package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetCardSuit(t *testing.T) {
	ace  := Kind{1, "Ace"}
	card := Card{Suit: Spades, Kind: ace }

	assert.Equal(t, Spades, card.Suit)
}

func TestSetCardKind(t *testing.T) {
	ace  := Kind{1, "Ace"}
	card := Card{Suit: Spades, Kind: ace }

	assert.Equal(t, ace, card.Kind)
}

func TestCardString(t *testing.T) {
	t.Error("Implement")
}
