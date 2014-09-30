package game

type Kind struct {
	Rank uint
	Name string
}

func (k Kind) Compare(other Kind) int {
	if k.Rank < other.Rank {
		return 1
	} else if k.Rank > other.Rank {
		return -1
	} else {
		return 0
	}
}
