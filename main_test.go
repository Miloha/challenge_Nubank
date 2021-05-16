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

	t.Run("General review of the function", func(t *testing.T) {
		testManagecard := managecard.NewCard()
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

	t.Run("Test card-not-active", func(t *testing.T) {
		testManagecard := managecard.NewCard()
		account := managecard.Account{
			ActiveCard:     false,
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

		got := reply.Violations[0]

		want := "card-not-active"

		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, reply)
		}
	})

	t.Run("Test account-not-initialized", func(t *testing.T) {
		testManagecard := managecard.NewCard()
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

		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		testManagecard.AddTransaction(transaction, &reply)

		got := reply.Violations[0]

		want := "account-not-initialized"

		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, reply)
		}
	})

	t.Run("Test insufficient-limit", func(t *testing.T) {
		testManagecard := managecard.NewCard()
		account := managecard.Account{
			ActiveCard:     true,
			AvailableLimit: 10,
		}

		reply := managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}

		transaction := managecard.Transaction{
			Merchant: "t",
			Amount:   100,
			Time:     "2020-01-02T15:04:05.000Z",
		}

		testManagecard.AddAccount(account, &reply)
		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		testManagecard.AddTransaction(transaction, &reply)

		got := reply.Violations[1]

		want := "insufficient-limit"

		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, reply)
		}
	})

	t.Run("Test doubled-transaction", func(t *testing.T) {
		testManagecard := managecard.NewCard()
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
		testManagecard.AddTransaction(transaction, &reply)
		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		testManagecard.AddTransaction(transaction, &reply)

		got := reply.Violations[1]

		want := "doubled-transaction"

		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, reply)
		}
	})

	t.Run("Test high-frequency-small-interval", func(t *testing.T) {
		testManagecard := managecard.NewCard()
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
		testManagecard.AddTransaction(transaction, &reply)
		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		transaction = managecard.Transaction{
			Merchant: "x",
			Amount:   10,
			Time:     "2020-01-02T15:04:05.000Z",
		}
		testManagecard.AddTransaction(transaction, &reply)
		reply = managecard.DataOuputs{
			Account:    account,
			Violations: []string{},
		}
		transaction = managecard.Transaction{
			Merchant: "y",
			Amount:   10,
			Time:     "2020-01-02T15:04:05.000Z",
		}
		testManagecard.AddTransaction(transaction, &reply)

		got := reply.Violations[1]

		want := "high-frequency-small-interval"

		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, reply)
		}
	})

}
