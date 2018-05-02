package main

import (
	"github.com/jimxshaw/Gopay/payment"
)

func main() {
	var option payment.PaymentOption

	option = payment.CreateCreditAccount(
		"James Madison",
		"1111-2222-3333-4444",
		7,
		2030,
		777)

	option.ProcessPayment(1000)

	option = payment.CreateCashAccount()

	option.ProcessPayment(800)
}
