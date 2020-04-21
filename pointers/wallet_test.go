package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(30))

	got := wallet.Balance()
	want:= Bitcoin(30)
	// fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
