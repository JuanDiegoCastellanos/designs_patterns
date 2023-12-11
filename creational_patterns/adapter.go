package main

import "fmt"

type Payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment Using Cash")
}
func ProcessPayment(p Payment) {
	p.Pay()
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

// BankPayment Adapter
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func mainx() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bpa := &BankPaymentAdapter{
		bankAccount: 22324,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)

}
