package domain

import (
	"fmt"
	"math/rand"
	"time"
)

type UnoAIPlayer struct {
	AbstractUnoPlayer
}

func NewUnoAIPlayer() (p *UnoAIPlayer) {
	UnoAIPlayer := &UnoAIPlayer{
		AbstractUnoPlayer: NewAbstractUnoPlayer(),
	}
	return UnoAIPlayer
}

func (p *UnoAIPlayer) NameSelf() string {
	rand.Seed(time.Now().UnixNano())
	p.name = fmt.Sprintf("AI Player %d", rand.Intn(999999))
	return p.name
}

func (p *UnoAIPlayer) Show() *UnoCard {
	rand.Seed(time.Now().UnixNano())
	var i int
	if len(p.hand) > 1 {
		i = rand.Intn(len(p.hand))
	}
	return &p.hand[i]
}
