package src

type Player interface {
	NameSelf()
	UseExchangeChance() bool
	ChoosePlayerForExchange() (player Player)
	Show() (card Card)
}

type AbstractPlayer struct {}

type PlayerAttribute struct {
	ExchangeChance int
	Points         int
	Name           string
}

func NewPlayerAttribute() (pa *PlayerAttribute) {
	return &PlayerAttribute{
		ExchangeChance: 1,
		Points:         0,
	}
}

func (p *AbstractPlayer) AddCardIntoHand(card Card) {
	//
}

func (p *AbstractPlayer) ExecuteExchangeHand() {
	//
}
