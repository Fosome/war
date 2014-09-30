package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHeartsSuit(t *testing.T) {
	assert.Equal(t, "Hearts", Hearts)
}

func TestDiamondsSuit(t *testing.T) {
	assert.Equal(t, "Diamonds", Diamonds)
}

func TestClubsSuit(t *testing.T) {
	assert.Equal(t, "Clubs", Clubs)
}

func TestSpadesSuit(t *testing.T) {
	assert.Equal(t, "Spades", Spades)
}
