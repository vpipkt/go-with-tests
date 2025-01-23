package wallet

type Mycoin int

type Wallet struct {
	balance Mycoin
}

func (w *Wallet) Deposit(amount Mycoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Mycoin {
	return w.balance
}
