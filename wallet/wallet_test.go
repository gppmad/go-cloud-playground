package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		deposit := Bitcoin(10.0)
		wallet.Deposit(deposit)

		assertBalance(t, wallet, 10.0)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20.0)}

		withdraw := Bitcoin(10.0)
		err := wallet.Withdraw(withdraw)

		if err != nil {
			t.Errorf("error is not nil but %s", err.Error())
		}

		assertBalance(t, wallet, 10.0)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20.0)}

		withdraw := Bitcoin(30.0)
		// In this case err is not nil.
		err := wallet.Withdraw(withdraw)

		assertError(t, err, ErrInsufficientAmount)
		// The balance should be the same
		assertBalance(t, wallet, 20.0)

	})

	t.Run("check error message", func(t *testing.T) {
		wallet := Wallet{amount: Bitcoin(20.0)}

		withdraw := Bitcoin(30.0)
		// In this case err is not nil.
		err := wallet.Withdraw(withdraw)

		assertError(t, err, ErrInsufficientAmount)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	balance := wallet.Balance()

	if balance != want {
		t.Errorf("Wallet Balance %s is different from the Deposit %s", balance, want)
	}
}

func assertError(t testing.TB, got, err error) {
	t.Helper()
	if err == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if err != got {
		t.Errorf("Error message is not accurate, got %s", err.Error())
	}
}
