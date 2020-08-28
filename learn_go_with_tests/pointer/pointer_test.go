package pointer

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()

	fmt.Printf("address of test is %v \n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
