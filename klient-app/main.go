package main

import (
	"fmt"
	"rajasureshaditya/go-workspace/constructinterface"
)

func main() {
	// var trans constructinterface.Paymethodinterface
	trans := &constructinterface.Creditcard{
		Name:          "Raja",
		Creditcard_no: "12345",
	}
	trans1 := constructinterface.NewPayment(trans)
	fmt.Println(trans1.Pay())
}
