package domain

import "fmt"

type UnoGame struct {
	AbstractGame[UnoCard, UnoPlayer]
	winner *UnoPlayer
	stack  []*UnoCard
}

func NewUnoGame() *UnoGame {
	unoGame := &UnoGame{}
	return unoGame
}

func (g *UnoGame) shouldBreakDrawCards(deck *Deck[UnoCard], count int) bool {
	return len(deck.getCards())-4*5 <= 0
}

func (g *UnoGame) takeRound() {

	fmt.Println("--------------------")
	fmt.Printf("Round %d \n", g.round)
	fmt.Println("--------------------")

	topCard := g.getTopCardFromStack()
	if topCard == nil {
		fmt.Println("there are no card in stack, draw a card from deck")
		topCard = g.safeDrawCard()
		g.stack = append(g.stack, topCard)
	}
	fmt.Printf("Top card color: %s, number: %s \n", topCard.Color, topCard.Number)

	for _, p := range g.players {
		fmt.Printf("It's %s's turn \n", p.GetName())

		//玩家沒有可出的牌
		if g.hasNoCardCanShow(p) {

			fmt.Println(123)
			// 玩家就必須從牌堆中抽一張牌，如果此時牌堆空了，則會先把檯面上除了最新的牌以外的牌放回牌堆中進行洗牌
			newCard := g.safeDrawCard()
			p.SetHand(append(p.GetHand(), *newCard))

		} else {
			card := p.Show()
			for !g.isValidateCard(card) {
				card = p.Show()
				fmt.Println(card)

			}

			removed := p.removeCardFromHand(*card)

			if len(p.GetHand()) == 0 {
				fmt.Printf("Player %s has no more hand.\n", p.GetName())
				g.winner = &p
				return
			} else {
				fmt.Printf("Player: %s shows card, color: %s, number: %s \n", p.GetName(), removed.Color, removed.Number)
				g.stack = append(g.stack, removed)
			}

		}
	}
	g.round += 1

}

func (g *UnoGame) end() {
	w := *g.winner
	fmt.Printf("The winner is %s \n", w.GetName())
}

func (g *UnoGame) hasNextRound() bool {
	return g.winner == nil
}

func (g *UnoGame) safeDrawCard() *UnoCard {
	newCard := g.deck.DrawCard()

	if newCard != nil {
		return newCard
	}
	g.deck.setCards(g.stack[:len(g.stack)-1])
	g.deck.Shuffle()

	return g.deck.DrawCard()
}

func (g *UnoGame) getTopCardFromStack() *UnoCard {
	if len(g.stack) == 0 {
		return nil
	}
	return g.stack[len(g.stack)-1]
}

func (g *UnoGame) isValidateCard(choseCard *UnoCard) bool {
	topCard := g.getTopCardFromStack()
	return choseCard.Color == topCard.Color || choseCard.Number == topCard.Number
}

func (g *UnoGame) hasNoCardCanShow(p UnoPlayer) bool {
	for _, card := range p.GetHand() {
		if g.isValidateCard(&card) {
			return false
		}
	}
	return true
}
