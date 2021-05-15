package main

import (
	"testing"

	"github.com/Miloha/challenge_Nubank/managecard"
)

func TestAddAccount(t *testing.T) {

	testManagecard := managecard.NewCard()

	t.Run("General review of the function", func(t *testing.T) {
		account := managecard.Account{
			ActiveCard:     false,
			AvailableLimit: 0,
		}

		reply := managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}

		got := testManagecard.AddAccount(account, &reply)
		var want error

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, account)
		}
	})

	t.Run("Approval account", func(t *testing.T) {
		account := managecard.Account{
			ActiveCard:     true,
			AvailableLimit: 0,
		}

		reply := managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}

		testManagecard.AddAccount(account, &reply)

		got := reply.Account.ActiveCard
		want := true

		if got != want {
			t.Errorf("got %t want %t given, %v", got, want, account)
		}
	})

}

func TestAddTransaction(t *testing.T) {

	testManagecard := managecard.NewCard()

	t.Run("General review of the function", func(t *testing.T) {
		transaction := managecard.Transaction{
			Merchant: "",
			Amount:   0,
			Time:     "",
		}

		account := managecard.Account{
			ActiveCard:     true,
			AvailableLimit: 0,
		}

		reply := managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}

		got := testManagecard.AddTransaction(transaction, &reply)
		var want error

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, account)
		}
	})

	t.Run("Approval transaction", func(t *testing.T) {
		account := managecard.Account{
			ActiveCard:     true,
			AvailableLimit: 100,
		}

		reply := managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}

		transaction := managecard.Transaction{
			Merchant: "t",
			Amount:   10,
			Time:     "2020-01-02T15:04:05.000Z",
		}

		testManagecard.AddAccount(account, &reply)
		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		testManagecard.AddTransaction(transaction, &reply)

		got := len(reply.Violations)

		want := 1

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, reply)
		}
	})

}
