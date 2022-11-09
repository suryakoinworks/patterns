package chainresponsibilities

import "patterns/observers"

const (
	BeforeRepayment = observers.Event("before_repayment")
	AfterRepayment  = observers.Event("after_repayment")
)

type (
	Payment struct {
		Method string
		Amount int
	}

	PaymentRule interface {
		Support(payment *Payment) bool
		Do(payment *Payment) error
	}

	Repayment struct {
		modificator  observers.Dispatcher
		paymentRules []PaymentRule
	}
)

func NewRepayment(modificator observers.Dispatcher, rules ...PaymentRule) Repayment {
	return Repayment{paymentRules: rules, modificator: modificator}
}

func (r *Repayment) Do(payment *Payment) error {
	// r.modificator.Dispatch(BeforeRepayment, payment)
	for _, rule := range r.paymentRules {
		if !rule.Support(payment) {
			continue
		}

		err := rule.Do(payment)
		if err != nil {
			return err
		}
	}

	return nil //r.modificator.Dispatch(AfterRepayment, payment)
}
