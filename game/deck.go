package game

type Deck interface {
	// Return the cards for the Deck
	Cards() []Card
}

type EmptyDeck struct{
	cards []Card
}

func NewEmptyDeck() *EmptyDeck {
	return &EmptyDeck{[]Card{}}
}

func (d *EmptyDeck) Cards() []Card {
	return d.cards
}

type StandardDeck struct{
	cards []Card
}

func NewStandardDeck() *StandardDeck {
	cards := []Card{
		Card{Spades, Kind{1, "Ace"}},
		Card{Spades, Kind{2, "King"}},
		Card{Spades, Kind{3, "Queen"}},
		Card{Spades, Kind{4, "Jack"}},
		Card{Spades, Kind{5, "10"}},
		Card{Spades, Kind{6, "9"}},
		Card{Spades, Kind{7, "8"}},
		Card{Spades, Kind{8, "7"}},
		Card{Spades, Kind{9, "6"}},
		Card{Spades, Kind{10, "5"}},
		Card{Spades, Kind{11, "4"}},
		Card{Spades, Kind{12, "3"}},
		Card{Spades, Kind{13, "2"}},
		Card{Clubs, Kind{1, "Ace"}},
		Card{Clubs, Kind{2, "King"}},
		Card{Clubs, Kind{3, "Queen"}},
		Card{Clubs, Kind{4, "Jack"}},
		Card{Clubs, Kind{5, "10"}},
		Card{Clubs, Kind{6, "9"}},
		Card{Clubs, Kind{7, "8"}},
		Card{Clubs, Kind{8, "7"}},
		Card{Clubs, Kind{9, "6"}},
		Card{Clubs, Kind{10, "5"}},
		Card{Clubs, Kind{11, "4"}},
		Card{Clubs, Kind{12, "3"}},
		Card{Clubs, Kind{13, "2"}},
		Card{Hearts, Kind{1, "Ace"}},
		Card{Hearts, Kind{2, "King"}},
		Card{Hearts, Kind{3, "Queen"}},
		Card{Hearts, Kind{4, "Jack"}},
		Card{Hearts, Kind{5, "10"}},
		Card{Hearts, Kind{6, "9"}},
		Card{Hearts, Kind{7, "8"}},
		Card{Hearts, Kind{8, "7"}},
		Card{Hearts, Kind{9, "6"}},
		Card{Hearts, Kind{10, "5"}},
		Card{Hearts, Kind{11, "4"}},
		Card{Hearts, Kind{12, "3"}},
		Card{Hearts, Kind{13, "2"}},
		Card{Diamonds, Kind{1, "Ace"}},
		Card{Diamonds, Kind{2, "King"}},
		Card{Diamonds, Kind{3, "Queen"}},
		Card{Diamonds, Kind{4, "Jack"}},
		Card{Diamonds, Kind{5, "10"}},
		Card{Diamonds, Kind{6, "9"}},
		Card{Diamonds, Kind{7, "8"}},
		Card{Diamonds, Kind{8, "7"}},
		Card{Diamonds, Kind{9, "6"}},
		Card{Diamonds, Kind{10, "5"}},
		Card{Diamonds, Kind{11, "4"}},
		Card{Diamonds, Kind{12, "3"}},
		Card{Diamonds, Kind{13, "2"}},
	}

	return &StandardDeck{cards}
}

func (d *StandardDeck) Cards() []Card {
	return d.cards
}

func (d *StandardDeck) Shuffle() {
	//
}
