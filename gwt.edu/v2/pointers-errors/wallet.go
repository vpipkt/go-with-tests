package wallet

import (
	"errors"
	"fmt"
)

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

var ErrInsufficientFunds = errors.New("User has insufficient funds for withdrawal requested")

func (w *Wallet) Withdraw(amount Mycoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
