package wallet

import (
	"fmt"
	"github.com/pkg/errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // same as dereferenced pointer (*w).balance
	/*
		Now you might wonder, why did they pass? We didn't dereference the pointer in the function, like so:
		and seemingly addressed the object directly. In fact, the code above using (*w) is absolutely valid.

		However, the makers of Go deemed this notation cumbersome,
		so the language permits us to write w.balance, without explicit dereference.

		These pointers to structs even have their own name: struct pointers and they are automatically dereferenced.
		https://golang.org/ref/spec#Method_values
	*/
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
