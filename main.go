package main

import (
	"bigger-or-smaller-game/domain"
)

func main() {
	game := domain.NewShowdownGame()

	game.Start()
}
