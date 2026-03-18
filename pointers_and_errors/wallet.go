package pointers_and_errors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds.")

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
func (w *Wallet) Deposit(amount BitCoin) {
	if amount < 0 {
		return
	}

	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.Balance() {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}
