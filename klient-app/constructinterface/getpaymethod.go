package constructinterface

import (
	"fmt"
)

type Paymethodinterface interface {
	Pay() error
}

type Creditcard struct {
	Name          string
	Creditcard_no string
}

type Bank struct {
}

type Payment struct {
	Paymethodinterface
}

type Paystruct struct {
	PaymentName string
}

func (creditcard Creditcard) Pay() error {
	fmt.Println("Payment selected is Creditcard" + creditcard.Name)
	return nil
}

func (banks *Bank) Pay() error {
	fmt.Println("Payment selected is Bank")
	return nil
}

func NewPayment(payinterface Paymethodinterface) *Payment {
	return &Payment{
		payinterface,
	}
}
