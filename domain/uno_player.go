package domain

type UnoPlayer interface {
	NameSelf() (name string)
	Show() (card *UnoCard)
	AddCardIntoHand(card UnoCard)
	SetName(string)
	GetName() (name string)
	SetGame(any)
	GetHand() []UnoCard
	SetHand([]UnoCard)
	hasNoCardCanShow(UnoCard) bool
}

type AbstractUnoPlayer struct {
	AbstractPlayer[UnoCard, UnoPlayer]
	NameSelf func() (name string)
	Show     func() (card *UnoCard)
}

func NewAbstractUnoPlayer() (pa AbstractUnoPlayer) {
	unoPlayer := AbstractUnoPlayer{
		AbstractPlayer: AbstractPlayer[UnoCard, UnoPlayer]{},
	}

	return unoPlayer
}

func (p *AbstractUnoPlayer) hasNoCardCanShow(topCard UnoCard) bool {
	var hasNoCardCanShow = true
	for _, card := range p.GetHand() {
		if card.Color == topCard.Color || card.Number == topCard.Number {
			return false
		}
	}
	return hasNoCardCanShow
}

func (p *AbstractUnoPlayer) isValidateCard(choseCard *UnoCard, topCard *UnoCard) bool {
	return choseCard.Color == topCard.Color || choseCard.Number == topCard.Number
}
