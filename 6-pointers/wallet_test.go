package pointers

import (
	"testing"
)

// make a Wallet struct which lets us deposit Bitcoin
func TestWallet(t *testing.T) {

	assertBalance := func(tb testing.TB, wallet Wallet, want Bitcoin) {
		tb.Helper()

		got := wallet.Balance()

		if got != want {
			tb.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(30))

		want := Bitcoin(70)

		assertBalance(t, wallet, want)
	})
}
