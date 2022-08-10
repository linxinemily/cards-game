package src

type ExchangeHand struct {
	shouldRollbackRound int
}

func NewExchangeHand(shouldRollbackRound int) (e *ExchangeHand) {
	return &ExchangeHand{
		shouldRollbackRound: shouldRollbackRound,
	}
}

func (e *ExchangeHand) Rollback() {
	//
}

func (e *ExchangeHand) Exchange(p1 Player, p2 Player) {
	//
}

func (e *ExchangeHand) haveToRollback() bool {
	return false
}
