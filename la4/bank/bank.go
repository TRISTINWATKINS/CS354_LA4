package bank

import (
	"fmt"
	"la4/account"
)

type Bank struct {
	Accounts map[string]account.Account
}

func New() *Bank {
	return &Bank{Accounts: make(map[string]account.Account)}
}

func (b *Bank) Add(a account.Account) {
	b.Accounts[a.GetNumber()] = a
}

func (b *Bank) Accrue(rate float64) {
	for _, a := range b.Accounts {
		a.Accrue(rate)
	}
}

func (b *Bank) AccrueConcurrent(rate float64) float64 {
	ch := make(chan float64)
	count := 0

	for _, a := range b.Accounts {
		count++
		go func(ac account.Account) {
			before := ac.Balance()
			ac.Accrue(rate)
			diff := ac.Balance() - before
			ch <- diff
		}(a)
	}

	var total float64
	for i := 0; i < count; i++ {
		total += <-ch
	}

	return total
}

func (b *Bank) String() string {
	r := ""
	for _, a := range b.Accounts {
		r += fmt.Sprintf("%s\n", a.String())
	}
	return r
}
