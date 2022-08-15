package main

import (
	"rajasureshaditya/go-workspace/payment-device/paymentInterface"
)

func main() {
	trans := &paymentInterface.Creditcard{
		Name:          "Raja",
		Creditcard_no: "12345",
	}
	// var f FinalDetails = values
	banktrans := &paymentInterface.Paypal{
		Name:  "Suresh",
		UPIid: "242526",
	}
	trans1 := paymentInterface.NewPayment(banktrans)
	trans1.Pay()
	var paymenttype paymentInterface.PaymentexpInterface = trans
	paymenttype.Pay()
	debitcard := &paymentInterface.DebitCard{
		Name:         "Raja",
		DebitCard_no: "12345",
	}
	paymenttype = debitcard
	paymenttype.Pay()
}
