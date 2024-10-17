package pointers

import (
	"testing"
)

// make a Wallet struct which lets us deposit Bitcoin
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(1))

	got := wallet.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
