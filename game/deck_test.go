package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEmptyDeck(t *testing.T) {
	deck := NewEmptyDeck()
	assert.Equal(t, 0, len(deck.Cards()))
}

func TestStandardDeck(t *testing.T) {
	deck := NewStandardDeck()
	assert.Equal(t, 52, len(deck.Cards()))
}
