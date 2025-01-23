package wallet

import "fmt"

type Mycoin int

func (m Mycoin) String() string {
	return fmt.Sprintf("%d MYC", m)
}

type Wallet struct {
	balance Mycoin
}

func (w *Wallet) Deposit(amount Mycoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Mycoin {
	return w.balance
}
