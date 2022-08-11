package domain

import (
	"bigger-or-smaller-game/domain/enum"
)

type Card struct {
	rank enum.Rank
	suit enum.Suit
}

func CompareTo(card Card) int {
	return 0
}
