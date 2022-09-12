package domain

import (
	"errors"
)

type ShowdownPlayer interface {
	NameSelf() (name string)
	ToUseExchangeChance() bool
	ChoosePlayerForExchange() (showdownPlayer ShowdownPlayer)
	Show() (card *ShowdownCard)
	AddCardIntoHand(card *ShowdownCard)
	CanUseExchangeHand() bool
	GetName() (name string)
	SetExchangeHand(*ExchangeHand)
	GetPoints() int
	SetPoints(int)
	SetGame(any)
	GetHand() []*ShowdownCard
	SetHand([]*ShowdownCard)
	GetExchangeHand() *ExchangeHand
}

type AbstractShowdownPlayer struct {
	AbstractPlayer[ShowdownCard, ShowdownPlayer]
	exchangeChance          int
	points                  int
	hand                    []*ShowdownCard
	exchangeHand            *ExchangeHand
	ChoosePlayerForExchange func() (showdownPlayer ShowdownPlayer)
	NameSelf                func() (name string)
	ToUseExchangeChance     func() bool
	Show                    func() (card *ShowdownCard)
}

func NewAbstractShowdownPlayer() (pa AbstractShowdownPlayer) {
	showdownPlayer := AbstractShowdownPlayer{
		AbstractPlayer: AbstractPlayer[ShowdownCard, ShowdownPlayer]{},
		exchangeChance: 1,
		points:         0,
	}

	return showdownPlayer
}

func (p *AbstractShowdownPlayer) AddCardIntoHand(card *ShowdownCard) {
	p.hand = append(p.hand, card)
}

// func (p *AbstractShowdownPlayer) GetName() string {
// 	return p.name
// }

func (p *AbstractShowdownPlayer) SetExchangeHand(exchangeHand *ExchangeHand) {
	p.exchangeHand = exchangeHand
	p.exchangeChance = p.exchangeChance - 1
	p.exchangeHand.Exchange()
}

func (p *AbstractShowdownPlayer) GetExchangeHand() *ExchangeHand {
	return p.exchangeHand
}

func (p *AbstractShowdownPlayer) CanUseExchangeHand() bool {
	return p.exchangeChance > 0
}

func (p *AbstractShowdownPlayer) removeCardFromHand(i int) (*ShowdownCard, error) {
	if i < 0 || i > len(p.hand)-1 {
		return nil, errors.New("invalid index")
	}
	card := p.hand[i]
	p.hand[i] = p.hand[len(p.hand)-1]
	p.hand = p.hand[:len(p.hand)-1]
	return card, nil
}

func (p *AbstractShowdownPlayer) GetPoints() int {
	return p.points
}

func (p *AbstractShowdownPlayer) SetPoints(points int) {
	p.points = points
}

func (p *AbstractShowdownPlayer) SetHand(cards []*ShowdownCard) {
	p.hand = cards
}

func (p *AbstractShowdownPlayer) GetHand() []*ShowdownCard {
	return p.hand
}
