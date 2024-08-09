package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(num Bitcoin) {
	w.balance += num
}

func (w *Wallet) Withdraw(num Bitcoin) error {
	if num > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= num
	return nil
}
