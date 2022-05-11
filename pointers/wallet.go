package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(quantity Bitcoin) {
	w.balance += quantity
}

var ErrInsufficientFunds = errors.New("Insufficient funds!")

func (w *Wallet) Withdraw(quantity Bitcoin) error {
	if w.balance >= quantity {
		w.balance -= quantity
		return nil
	}
	return ErrInsufficientFunds

}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
