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
		wallet := Wallet{BitCoin(10)}
		wallet.Withdraw(BitCoin(1))

		want := BitCoin(9)

		assert(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{BitCoin(10)}
		err := wallet.Withdraw(BitCoin(11))
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertError(t *testing.T, err error, want error) {
	t.Helper()
	if err == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if err != want {
		t.Errorf("got %q, want %q", err, want)
	}
}
