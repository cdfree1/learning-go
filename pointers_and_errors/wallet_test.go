package pointers_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("Got %s but wanted %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one.")
		}

		if got.Error() != ErrInsufficientFunds.Error() {
			t.Errorf("got %q but wanted %q", got, ErrInsufficientFunds)
		}

	}

	assertNoError := func(t testing.TB, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error but didn't want one!")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: 9}

		wallet.Deposit(10)

		assertBalance(t, wallet, BitCoin(19))
	})

	t.Run("withdraw sufficient amount", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}

		err := wallet.Withdraw(BitCoin(5))

		assertNoError(t, err)

		assertBalance(t, wallet, BitCoin(5))
	})
	t.Run("withdraw insufficient amount", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(10)}

		err := wallet.Withdraw(BitCoin(100))

		assertError(t, err)

		assertBalance(t, wallet, BitCoin(10))

	})
}
