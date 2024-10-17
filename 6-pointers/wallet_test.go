package pointers

import (
	"errors"
	"testing"
)

var WithdrawErrorMsg = errors.New("cannot withdraw, insufficient funds")

// make a Wallet struct which lets us deposit Bitcoin
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(110))

		assertError(t, err, WithdrawErrorMsg)
		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingAmount := Bitcoin(20)
		wallet := Wallet{balance: startingAmount}
		err := wallet.Withdraw(Bitcoin(110))

		assertError(t, err, WithdrawErrorMsg)
		assertBalance(t, wallet, startingAmount)
	})
}

func assertBalance(tb testing.TB, wallet Wallet, want Bitcoin) {
	tb.Helper()

	got := wallet.Balance()

	if got != want {
		tb.Errorf("got %s want %s", got, want)
	}
}

func assertError(tb testing.TB, got, want error) {
	tb.Helper()

	if got == nil {
		tb.Fatal("didnt get an error but wanted one")
	}

	if !errors.Is(got, want) {
		tb.Errorf("got %q want %q", got, want)
	}
}
