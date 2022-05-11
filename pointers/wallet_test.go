package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	checkBalance := func(t testing.TB, got, want fmt.Stringer) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)
		checkBalance(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Withdraw(5)

		got := wallet.Balance()
		want := Bitcoin(5)
		checkBalance(t, got, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		wallet.Withdraw(20)

		got := wallet.Balance()
		want := Bitcoin(10)
		checkBalance(t, got, want)
	})
}
