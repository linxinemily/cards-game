package enum

type Suit int

const (
	Club Suit = iota
	Diamond
	Heart
	Spade
)

func (s Suit) String() string {
	switch s {
	case Club:
		return "Club"
	case Diamond:
		return "Diamond"
	case Heart:
		return "Heart"
	case Spade:
		return "Spade"
	default:
		return "Unknown"
	}
}
