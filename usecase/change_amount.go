package usecase

import (
	"patterns/chainresponsibilities"
	"patterns/observers"
)

type ChangeAmount struct {
}

func (c *ChangeAmount) Handle(payload interface{}) {
	payment, ok := payload.(*chainresponsibilities.Payment)
	if !ok {
		return
	}

	if payment.Method == InstaMoneyMethod {
		return
	}

	payment.Amount = 1000
}

func (c *ChangeAmount) Listen() observers.Event {
	return chainresponsibilities.BeforeRepayment
}

func (c *ChangeAmount) Priority() int {
	return 0
}
