package checkingaccount

import (
	"fmt"
	"la4/account"
	"la4/customer"
)

type CheckingAccount struct {
	Number   string
	Customer *customer.Customer
	BalanceV float64
}

func New(number string, c *customer.Customer, bal float64) *CheckingAccount {
	return &CheckingAccount{
		Number:   number,
		Customer: c,
		BalanceV: bal,
	}
}

func (a *CheckingAccount) Accrue(rate float64) {
}

func (a *CheckingAccount) Balance() float64 {
	return a.BalanceV
}

func (a *CheckingAccount) Deposit(amount float64) {
	a.BalanceV += amount
}

func (a *CheckingAccount) Withdraw(amount float64) {
	a.BalanceV -= amount
}

func (a *CheckingAccount) GetNumber() string {
	return a.Number
}

func (a *CheckingAccount) GetCustomer() *customer.Customer {
	return a.Customer
}

func (a *CheckingAccount) String() string {
	return fmt.Sprintf("%s:%s:%.2f",
		a.Number, a.Customer.String(), a.BalanceV)
}

var _ account.Account = (*CheckingAccount)(nil)
