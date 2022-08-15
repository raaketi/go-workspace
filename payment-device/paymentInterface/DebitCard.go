package paymentInterface

import (
	"fmt"
)

type DebitCard struct {
	Name         string
	DebitCard_no string
}

func (debitcard DebitCard) Pay() error {
	fmt.Println("Payment selected is DebitCard" + debitcard.Name)
	return nil
}
