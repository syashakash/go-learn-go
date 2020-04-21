package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(30)

	got := wallet.Balance()
	want:= 30
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
