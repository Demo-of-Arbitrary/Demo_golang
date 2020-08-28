package pointer

import "fmt"

type BitCoin int

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance BitCoin
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Deposit(amount BitCoin) {
	w.balance += amount
	fmt.Printf("address of package is %v \n", &w.balance)
}

func (w *Wallet) Withdraw(amount BitCoin) {
	w.balance -= amount
}
