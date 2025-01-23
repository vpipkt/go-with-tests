package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	wallet := Wallet{}
	wallet.Deposit(Mycoin(10))

	got := wallet.Balance()

	want := Mycoin(10)

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
