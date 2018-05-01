package payment

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// PaymentOption any interaction with accounts must go through this interface.
type PaymentOption interface {
	ProcessPayment(float32) bool
}

// CreditCard data structure
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// CreateCreditAccount constructor function
func CreateCreditAccount(ownerName, cardNumber string, expirationMonth, expirationYear, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// ProcessPayment takes in the amount that's being processed and return a boolean depending on whether or not the processing was successful.
func (c *CreditCard) ProcessPayment(amount float32) bool {
	fmt.Println("Processing a credit card payment...")
	return true
}

// OwnerName getter
func (c CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName setter
func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid owner name provided")
	}

	c.ownerName = value
	return nil
}

// CardNumber getter
func (c CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber setter
func (c *CreditCard) SetCardNumber(value string) error {
	if len(value) == 0 {
		return errors.New("Card number cannot be blank")
	}

	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid credit card number format")
	}

	c.cardNumber = value
	return nil
}

// ExpirationDate getter
func (c CreditCard) ExpirationDate() (int, int) {
	return c.expirationMonth, c.expirationYear
}

// SetExpirationDate setter
func (c *CreditCard) SetExpirationDate(month, year int) error {
	now := time.Now()

	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date must be in the future")
	}

	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode getter
func (c CreditCard) SecurityCode() int {
	return c.securityCode
}

// SetSecurityCode setter
func (c *CreditCard) SetSecurityCode(value int) error {
	if value < 100 || value > 999 {
		return errors.New("Invalid security code format")
	}

	c.securityCode = value
	return nil
}

// AvailableCredit getter
func (c CreditCard) AvailableCredit() float32 {
	// This value can be retrieved from a web service, client
	// doesn't need to know where it's from.
	return 7000
}
