package domain

type UnoPlayer interface {
	Player[UnoCard]
}

type AbstractUnoPlayer struct {
	AbstractPlayer[UnoCard, UnoPlayer]
}

func NewAbstractUnoPlayer() (pa AbstractUnoPlayer) {
	unoPlayer := AbstractUnoPlayer{
		AbstractPlayer: AbstractPlayer[UnoCard, UnoPlayer]{},
	}

	return unoPlayer
}
