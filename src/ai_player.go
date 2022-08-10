package src

type AIPlayer struct {
	AbstractPlayer
	*PlayerAttribute
}

func NewAIPlayer() (p *AIPlayer) {
	return &AIPlayer{
		PlayerAttribute: NewPlayerAttribute(),
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
