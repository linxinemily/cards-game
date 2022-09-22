package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type UnoHumanPlayer struct {
	AbstractUnoPlayer
}

func NewUnoHumanPlayer() *UnoHumanPlayer {
	return &UnoHumanPlayer{
		AbstractUnoPlayer: NewAbstractUnoPlayer(),
	}
}

func (p *UnoHumanPlayer) NameSelf() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter player name:")
	scanner.Scan()
	name := scanner.Text()
	p.name = name
	return name
}

func (p *UnoHumanPlayer) Show() (card *UnoCard) {

	var choseCard *UnoCard

	for choseCard == nil {

		for i, c := range p.hand {
			fmt.Printf("[%d] color %s, number %s \n", i, c.Color, c.Number)
		}

		fmt.Println("Choose a card to show, enter index of the card:")
		fmt.Println("Note: The Chose card should be same color or number with the card on top of the stack.")

		var intVar int
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		intVar, err := strconv.Atoi(scanner.Text())

		choseCard, err = p.getCardFromHand(intVar)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
	}

	return choseCard
}
