package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	amount Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.amount
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.amount += amount
}

var ErrInsufficientAmount = errors.New("you are withdrawing too much")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.amount {
		return ErrInsufficientAmount
	}
	w.amount -= amount
	return nil

}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
