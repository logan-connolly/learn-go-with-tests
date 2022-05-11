package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		tenBitcoin := Bitcoin(10)
		wallet.Deposit(tenBitcoin)
		assertBalance(t, wallet, tenBitcoin)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(5)
		want := Bitcoin(5)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(20)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w Wallet, want fmt.Stringer) {
	t.Helper()
	if want != w.Balance() {
		t.Errorf("got %s want %s", w.Balance(), want)
	}
}
func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("Did not expect to get an error")
	}
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Error("Wanted an error, but didn't get one")
	}
	if got != want {
		t.Errorf("expected %q got %q", want, got)
	}
}
