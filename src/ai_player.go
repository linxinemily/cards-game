package src

type AIPlayer struct {
	AbstractPlayer
	playerAttribute *PlayerAttribute
}

func NewAIPlayer() (p *AIPlayer) {
	return &AIPlayer{
		playerAttribute: NewPlayerAttribute(),
	}
}

func (p *AIPlayer) NameSelf() {

}

func (p *AIPlayer) UseExchangeChance() bool {
	return false
}

func (p *AIPlayer) ChoosePlayerForExchange() (player Player) {
	return nil
}

func (p *AIPlayer) Show() (card Card) {
	return Card{}
}
