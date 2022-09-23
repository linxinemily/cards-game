package main

import (
	"bigger-or-smaller-game/domain"
)

func main() {
	/** Showdown Game **/
	showdownGame := domain.IGame[domain.ShowdownCard, domain.ShowdownPlayer]{
		Game: domain.NewShowdownGame(),
	}
	showdownGame.InitPlayers([]domain.ShowdownPlayer{
		domain.NewShowdownAIPlayer(),
		domain.NewShowdownAIPlayer(),
		domain.NewShowdownAIPlayer(),
		domain.NewShowdownHumanPlayer(),
	})

	showdownGame.Start(*domain.NewShowdownDeck())

	/** Uno Game **/
	UnoGame := domain.IGame[domain.UnoCard, domain.UnoPlayer]{
		Game: domain.NewUnoGame(),
	}
	UnoGame.InitPlayers([]domain.UnoPlayer{
		domain.NewUnoAIPlayer(),
		domain.NewUnoAIPlayer(),
		domain.NewUnoAIPlayer(),
		domain.NewUnoHumanPlayer(),
	})

	UnoGame.Start(*domain.NewUnoDeck())
}
