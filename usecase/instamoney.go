package usecase

import (
	"fmt"
	"patterns/chainresponsibilities"
	"strings"
)

const InstaMoneyMethod string = "INSTA_MONEY"

type InstaMoneyRule struct {
}

func (n *InstaMoneyRule) Support(payment *chainresponsibilities.Payment) bool {
	return strings.ToUpper(payment.Method) == InstaMoneyMethod
}

func (n *InstaMoneyRule) Do(payment *chainresponsibilities.Payment) error {
	fmt.Println(payment.Method)
	fmt.Println(payment.Amount)

	return nil
}
