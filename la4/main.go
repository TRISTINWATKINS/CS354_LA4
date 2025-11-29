package main

import (
	"fmt"
	"la4/bank"
	"la4/checkingaccount"
	"la4/customer"
	"la4/savingaccount"
)

func main() {
	b := bank.New()
	c := customer.New("Ann")

	b.Add(checkingaccount.New("01001", c, 100.0))
	b.Add(savingaccount.New("01002", c, 200.0))

	fmt.Println("Before accrue:")
	fmt.Println(b.String())

	total := b.AccrueConcurrent(0.02)
	fmt.Println("Total interest earned:", total)

	fmt.Println("After accrue:")
	fmt.Println(b.String())
}
