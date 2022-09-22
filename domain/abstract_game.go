package domain

import (
	"fmt"
)

type IGame[C Card[C], P Player[C]] struct {
	Game[C, P]
}

func (g *IGame[C, P]) Start(deck Deck[C]) {

	g.setDeck(deck)

	g.init()

	g.playerDrawCards()

	for {
		if !g.hasNextRound() {
			break
		}
		g.takeRound()
	}

	g.end()
}

func (g *IGame[C, P]) playerDrawCards() {
	count := 0
	for {
		if g.shouldBreakDrawCards(g.getDeck(), count) {
			break
		}
		deck := g.getDeck()
		card := deck.DrawCard()
		p := g.getPlayers()[count%4]
		p.AddCardIntoHand(*card)
		count += 1
	}

}

type Game[C Card[C], P Player[C]] interface {
	init()
	hasNextRound() bool
	takeRound()
	end()
	setDeck(Deck[C])
	getDeck() *Deck[C]
	InitPlayers([]P)
	getPlayers() []P
	shouldBreakDrawCards(deck *Deck[C], count int) bool
}

func NewAbstractGame[C Card[C], P Player[C]]() *AbstractGame[C, P] {
	abstractGame := &AbstractGame[C, P]{}
	abstractGame.round = 1

	return abstractGame
}

type AbstractGame[C Card[C], P Player[C]] struct {
	round   int
	deck    *Deck[C]
	players []P
}

func (g *AbstractGame[C, P]) init() {
	for _, p := range g.getPlayers() {
		name := p.NameSelf()
		fmt.Printf("Player %s is added. \n", name)

		p.SetGame(g)
	}

	d := g.getDeck()
	d.Shuffle()
}

func (g *AbstractGame[C, P]) getPlayers() []P {
	return g.players
}

func (g *AbstractGame[C, P]) shouldBreakDrawCards(deck *Deck[C], count int) bool {
	return len(deck.getCards()) == 0
}

func (g *AbstractGame[C, P]) getDeck() *Deck[C] {
	return g.deck
}

func (g *AbstractGame[C, P]) setDeck(deck Deck[C]) {
	g.deck = &deck
}

func (g *AbstractGame[C, P]) InitPlayers(players []P) {
	g.players = players
}
