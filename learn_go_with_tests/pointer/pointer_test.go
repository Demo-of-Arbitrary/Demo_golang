package pointer

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t *testing.T, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))

		fmt.Printf("address of test is %v \n", &wallet.balance)

		want := BitCoin(10)
		assert(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		wallet.Withdraw(BitCoin(1))

		want := BitCoin(9)

		assert(t, wallet, want)
	})
}
