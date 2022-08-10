package src

import (
	"bigger-or-smaller-game/src/enum"
)

type Card struct {
	rank enum.Rank
	suit enum.Suit
}

func CompareTo(card Card) int {
	return 0
}
