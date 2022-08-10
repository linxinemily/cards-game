package src

type Player interface {
	NameSelf()
	UseExchangeChance() bool
	ChoosePlayerForExchange() (player Player)
	Show() (card Card)
}

type AbstractPlayer struct{}

type PlayerAttribute struct {
	exchangeChance int
	points         int
	name           string
	game           Game
	hand           []*Card
	exchangeHand   ExchangeHand
}

func NewPlayerAttribute() (pa *PlayerAttribute) {
	return &PlayerAttribute{
		exchangeChance: 1,
		points:         0,
	}
}

func (p *AbstractPlayer) AddCardIntoHand(card Card) {
	//
}

func (p *AbstractPlayer) ExecuteExchangeHand() {
	//
}
