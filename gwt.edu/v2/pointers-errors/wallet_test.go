package wallet

import (
	"fmt"
	"testing"
)

func assertBalance(t testing.TB, wallet Wallet, want Mycoin) {
	t.Helper()

	got := wallet.Balance()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		// Fatal to break out of here
		t.Fatal("expected error on insufficient fund, no error")
	}
	if got != want {
		t.Errorf("wanted error '%v' but got '%v'", want, got)
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit into new", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Mycoin(10))

		assertBalance(t, wallet, Mycoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Mycoin(42)}
		err := wallet.Withdraw(Mycoin(12))
		assertBalance(t, wallet, Mycoin(30))
		if err != nil {
			t.Errorf("got an error '%v' but didnt want one", err)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		start := Mycoin(100)
		wallet := Wallet{balance: start}
		err := wallet.Withdraw(Mycoin(200))
		assertBalance(t, wallet, start)
		assertError(t, err, ErrInsufficientFunds)
	})

	t.Run("stringer ok", func(t *testing.T) {
		v := Mycoin(123)
		got := fmt.Sprintf("%v", v)
		want := "123 MYC"
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
