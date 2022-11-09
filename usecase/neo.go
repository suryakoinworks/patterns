package usecase

import (
	"fmt"
	"patterns/chainresponsibilities"
	"strings"
)

const NeoMethod string = "NEO"

type NeoRule struct {
}

func (n *NeoRule) Support(payment *chainresponsibilities.Payment) bool {
	return strings.ToUpper(payment.Method) == NeoMethod
}

func (n *NeoRule) Do(payment *chainresponsibilities.Payment) error {
	fmt.Println(payment.Method)
	fmt.Println(payment.Amount)

	return nil
}
