package main

import (
	"bigger-or-smaller-game/domain"
	"fmt"
)

func main() {
	fmt.Println("🃏 Bigger or Smaller Game")

	game := domain.NewGame()

	game.Start()

	game.PlayerDrawCards()

	for {
		if !game.HasNextRound() {
			break
		}
		game.TakeRound()
	}

	game.End()
}
