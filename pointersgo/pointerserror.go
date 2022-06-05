package pointersgo

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Dollar int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (d Dollar) StringDollar() string {
	return fmt.Sprintf("Текущий курс BTC %d USD", d)
}

type WalletDol struct {
	balance Dollar
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

//dollar
func (d *WalletDol) Insert(amount Dollar) {
	d.balance = amount
}

func (d *WalletDol) BalanceDollar() Dollar {
	return d.balance
}

//exchange
func (d *WalletDol) Change(w *Wallet) int {
	return int(w.balance)*int(d.balance)
}