package domain

import (
	"bigger-or-smaller-game/domain/enum"
)

type ShowdownDeck struct {
	cards []*ShowdownCard
}

func NewShowdownDeck() (d *Deck[ShowdownCard]) {
	cards := make([]*ShowdownCard, 52)
	var count int
	for suit := enum.Suit(0); suit < enum.Spade+1; suit++ {
		for rank := enum.Rank(0); rank < enum.A+1; rank++ {
			cards[count] = NewCard(rank, suit)
			count++
		}
	}

	return &Deck[ShowdownCard]{
		cards: cards,
	}
}
