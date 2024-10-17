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
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		wallet.Withdraw(Bitcoin(110))
		assertBalance(t, wallet, Bitcoin(100))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingAmount := Bitcoin(100)
		wallet := Wallet{balance: startingAmount}
		err := wallet.Withdraw(Bitcoin(110))

		assertBalance(t, wallet, startingAmount)

		if err == nil {
			t.Error("wanted an error but didnt get one")
		}
	})
}
