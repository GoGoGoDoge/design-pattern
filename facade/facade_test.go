package facade_test

import (
	"facade"
	"testing"
)

func TestFacade(t *testing.T) {
	wallet := facade.NewWalletFacade()
	if wallet.AddMoney(10) != "Email: 10 has been saved to your account" {
		t.Fatal("Expecting", wallet.AddMoney(10))
	}
}
