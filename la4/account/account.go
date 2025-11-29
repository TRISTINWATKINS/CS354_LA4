package account

import "la4/customer"

type Account interface {
	Accrue(rate float64)
	Balance() float64
	Deposit(amount float64)
	Withdraw(amount float64)
	String() string
	GetNumber() string
	GetCustomer() *customer.Customer
}
