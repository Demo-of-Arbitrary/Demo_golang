package pointer

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(BitCoin(10))
	got := wallet.Balance()

	fmt.Printf("address of test is %v \n", &wallet.balance)

	want := BitCoin(10)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
