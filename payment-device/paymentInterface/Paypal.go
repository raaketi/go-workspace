package paymentInterface

import "fmt"

type Paypal struct {
	Name  string
	UPIid string
}

func (paypal *Paypal) Pay() error {
	fmt.Println("payment selected Paypal system")
	return nil
}
