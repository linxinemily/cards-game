package domain

type Game struct {
	round   int
	deck    Deck
	players []*Player
}

func (g *Game) Start() {
	//
}

func (g *Game) PlayerDrawCards() {
	//
}

func (g *Game) TakeRound() {
	//
}

func (g *Game) End() {
	//
}
