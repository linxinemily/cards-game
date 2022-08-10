package src

type HumanPlayer struct {
	AbstractPlayer
	playerAttribute *PlayerAttribute
}

func NewHumanPlayer() (p *HumanPlayer) {
	return &HumanPlayer{
		playerAttribute: NewPlayerAttribute(),
	}
}

func (p *HumanPlayer) NameSelf() {

}

func (p *HumanPlayer) UseExchangeChance() bool {
	return false
}

func (p *HumanPlayer) ChoosePlayerForExchange() (player Player) {
	return nil
}

func (p *HumanPlayer) Show() (card Card) {
	return Card{}
}
