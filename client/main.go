package main

import "fmt"

//"github.com/jimxshaw/Gopay/payment"

type CreditAccount struct{}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment...")
}

func CreateCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccout := &CreditAccount{}

	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccout.processPayment(amount)
		}
	}(chargeCh)

	return creditAccout
}

func main() {
	// var option payment.PaymentOption

	// option = payment.CreateCreditAccount(
	// 	"James Madison",
	// 	"1111-2222-3333-4444",
	// 	7,
	// 	2030,
	// 	777)

	// option.ProcessPayment(1000)

	// option = payment.CreateCashAccount()

	// option.ProcessPayment(800)

	/* MESSAGE PASSING */
	chargeCh := make(chan float32)

	CreateCreditAccount(chargeCh)

	chargeCh <- 1000

	// Prevent the main function from shutting down.
	var a string
	fmt.Scanln(&a)
}
