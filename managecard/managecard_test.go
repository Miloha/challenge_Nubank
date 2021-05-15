package managecard

import (
	"testing"

	"github.com/Miloha/challenge_Nubank/managecard"
)

func TestAddAccount(t *testing.T) {

	testManagecard := NewCard()
	account := managecard.Account{
		ActiveCard:     false,
		AvailableLimit: 0,
	}

	reply := managecard.DataOuputs{
		Account:    account,
		Violations: "",
	}

	got := testManagecard.AddAccount(account, reply)

	var want managecard.Ouputs

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, account)
	}

}
