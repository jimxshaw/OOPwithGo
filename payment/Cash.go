package payment

import (
	"fmt"
)

// Cash data structure.
type Cash struct{}

// CreateCashAccount constructor function.
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment takes in the amount to be processed and returns a boolean depending on if the processing is successful.
func (c Cash) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a cash transaction...")
	return true
}
