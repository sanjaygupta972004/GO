package main

import (
	"fmt"
)

type creditCard struct {
	cardNumber string
}

type googlePay struct {
	upiNumber string
}

func (cc creditCard) Pay(amount float64) {
	fmt.Printf("paid %.2f using credit card {card number %s}\n", amount, cc.cardNumber)
}

func (gp googlePay) Pay(amount float64) {
	fmt.Printf("paid %.2f using google pay {upi number %s}\n", amount, gp.upiNumber)
}

// illustrate this things how can i use interface
type PaymentMethod interface {
	Pay(amount float64)
}

type product struct {
	id          uint
	price       string
	name        string
	description string
	imageURL    string
}

type cart struct {
	item       []product
	totalPrice float64
}

func ProcessPayment(pm PaymentMethod, amount float64) {
	pm.Pay(amount)
}

func main() {
	fmt.Println("LEARNING ABOUT METHODS AND STRUCT IN GO LANG")

	creditCard := creditCard{cardNumber: "453464-5-4534-243"}
	googlePay := googlePay{upiNumber: "34cj49596ur38947rvd"}

	ProcessPayment(creditCard, 999)
	ProcessPayment(googlePay, 4444)

	// example 1
	err := Exa1()
	if err != nil {
		fmt.Printf("error in example 1 : %s\\n", err)
		return
	}

}
