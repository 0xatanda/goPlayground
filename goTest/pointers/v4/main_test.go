package main

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("Wanted an error but didn't get one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withrawal", func(t *testing.T) {
		wallt := Wallet{balance: Bitcoin(20)}
		wallt.Withdraw(Bitcoin(10))
		assertBalance(t, wallt, Bitcoin(10))
	})

	t.Run("Withdraw insuffience fund", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(1000))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}
