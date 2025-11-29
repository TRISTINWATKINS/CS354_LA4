package savingaccount

import (
	"fmt"
	"la4/account"
	"la4/customer"
)

type SavingAccount struct {
	Number   string
	Customer *customer.Customer
	BalanceV float64
	Interest float64
}

func New(number string, c *customer.Customer, bal float64) *SavingAccount {
	return &SavingAccount{
		Number:   number,
		Customer: c,
		BalanceV: bal,
	}
}

func (a *SavingAccount) Accrue(rate float64) {
	earned := a.BalanceV * rate
	a.Interest += earned
	a.BalanceV += earned
}

func (a *SavingAccount) Balance() float64 {
	return a.BalanceV
}

func (a *SavingAccount) Deposit(amount float64) {
	a.BalanceV += amount
}

func (a *SavingAccount) Withdraw(amount float64) {
	a.BalanceV -= amount
}

func (a *SavingAccount) GetNumber() string {
	return a.Number
}

func (a *SavingAccount) GetCustomer() *customer.Customer {
	return a.Customer
}

func (a *SavingAccount) String() string {
	return fmt.Sprintf("%s:%s:%.2f",
		a.Number, a.Customer.String(), a.BalanceV)
}

var _ account.Account = (*SavingAccount)(nil)
