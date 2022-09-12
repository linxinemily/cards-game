package domain

type Player[C Card[C]] interface {
	NameSelf() (name string)
	Show() (card *C)
	AddCardIntoHand(card *C)
	GetName() (name string)
	GetHand() []*C
	SetHand([]*C)
	SetGame(any)
}

type AbstractPlayer[C Card[C], P Player[C]] struct {
	name     string
	hand     []*Card[C]
	NameSelf func() string
	game     *AbstractGame[C, P]
}

func (p *AbstractPlayer[C, P]) GetName() string {
	return p.name
}

func (p *AbstractPlayer[C, P]) SetGame(game any) {

	other := game.(*AbstractGame[C, P])

	p.game = other

}
