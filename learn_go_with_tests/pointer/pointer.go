package pointer

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
	fmt.Printf("address of package is %v \n", &w.balance)
}
