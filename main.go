package main

import (
	"fmt"
	"patterns/chainresponsibilities"
	"patterns/observers"
	"patterns/usecase"
)

func main() {
	modificator := observers.Dispatcher{}

	changeAmount := usecase.ChangeAmount{}
	changeMethod := usecase.ChangeMethodForInstaMoney{}

	modificator.Register(&changeAmount, &changeMethod)

	repayment := chainresponsibilities.NewRepayment(
		modificator,
		&usecase.NeoRule{},
		&usecase.InstaMoneyRule{},
	)

	payment := chainresponsibilities.Payment{Method: usecase.InstaMoneyMethod, Amount: 10000}

	repayment.Do(&payment)

	fmt.Println("============= AFTER REPAYMENT ===============")
	fmt.Println(payment.Method)
	fmt.Println(payment.Amount)
}
