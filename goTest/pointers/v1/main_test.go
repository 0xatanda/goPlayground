package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	fmt.Printf("addresss of balancein test is %p \n", &wallet.balance)

	want := 10

	if got != want {
		t.Errorf("got %d but want %d", got, want)
	}
}
