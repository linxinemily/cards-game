package main

import (
	"bigger-or-smaller-game/domain"
)

func main() {
	game := domain.IGame[domain.UnoCard, domain.UnoPlayer]{
		Game: domain.NewUnoGame(),
	}
	game.InitPlayers([]domain.UnoPlayer{
		domain.NewUnoAIPlayer(),
		domain.NewUnoAIPlayer(),
		domain.NewUnoAIPlayer(),
		domain.NewUnoHumanPlayer(),
	})

	game.Start(*domain.NewUnoDeck())
}
