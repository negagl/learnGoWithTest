package pointers

import (
	"testing"
)

// make a Wallet struct which lets us deposit Bitcoin
func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(30))

		got := wallet.Balance()
		want := Bitcoin(70)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
