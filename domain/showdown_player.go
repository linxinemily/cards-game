package domain

type ShowdownPlayer interface {
	NameSelf() (name string)
	ToUseExchangeChance() bool
	ChoosePlayerForExchange() (showdownPlayer ShowdownPlayer)
	Show() (card *ShowdownCard)
	AddCardIntoHand(card ShowdownCard)
	CanUseExchangeHand() bool
	GetName() (name string)
	SetName(string)
	SetExchangeHand(*ExchangeHand)
	GetPoints() int
	SetPoints(int)
	SetGame(any)
	GetHand() []ShowdownCard
	SetHand([]ShowdownCard)
	GetExchangeHand() *ExchangeHand
	removeCardFromHand(i int) (*ShowdownCard, error)
}

type AbstractShowdownPlayer struct {
	AbstractPlayer[ShowdownCard, ShowdownPlayer]
	exchangeChance          int
	points                  int
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

func (p *AbstractShowdownPlayer) GetPoints() int {
	return p.points
}

func (p *AbstractShowdownPlayer) SetPoints(points int) {
	p.points = points
}
