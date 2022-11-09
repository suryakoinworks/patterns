package usecase

import (
	"patterns/chainresponsibilities"
	"patterns/observers"
)

type ChangeMethodForInstaMoney struct {
}

func (c *ChangeMethodForInstaMoney) Handle(payload interface{}) {
	payment, ok := payload.(*chainresponsibilities.Payment)
	if !ok {
		return
	}

	if payment.Method != InstaMoneyMethod {
		return
	}

	payment.Method = "OY_METHOD"
}

func (c *ChangeMethodForInstaMoney) Listen() observers.Event {
	return chainresponsibilities.AfterRepayment
}

func (c *ChangeMethodForInstaMoney) Priority() int {
	return 0
}
