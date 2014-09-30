package game

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSetKindRank(t *testing.T) {
	kind := Kind{Rank: 1, Name: "Ace"}
	assert.Equal(t, 1, kind.Rank)
}

func TestSetKindName(t *testing.T) {
	kind := Kind{Rank: 1, Name: "Ace"}
	assert.Equal(t, "Ace", kind.Name)
}

func TestComparingKinds(t *testing.T) {
	kind := Kind{Rank: 2, Name: "King"}

	// Return 0 when equal
	assert.Equal(t, 0, kind.Compare(Kind{2, "King"}))

	// Return 1 when greater
	assert.Equal(t, 1, kind.Compare(Kind{3, "Queen"}))

	// Return -1 when greater
	assert.Equal(t, -1, kind.Compare(Kind{1, "Ace"}))
}
