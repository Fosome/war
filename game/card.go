package game

import "fmt"

type Card struct {
	Suit Suit
	Kind Kind
}

func (c Card) String() string {
	return fmt.Sprintf("%s-%s", c.Kind.Name, c.Suit)
}
