package managecard

import (
	"fmt"
	"strings"
	"time"
)

//Account struct represents a card
type Account struct {
	ActiveCard     bool
	AvailableLimit float64
}

//Transaction struct represents a Ttansaction
type Transaction struct {
	Merchant string
	Amount   float64
	Time     string
}

// Account DB struct represents a account.
type Ouputs struct {
	database map[string]interface{} // private

}

// Reply struct.
type DataOuputs struct {
	Account    Account
	Violations []string
}

/*---------------*/

// Add methods adds an account to the struct (procedure).
func (b *Ouputs) AddAccount(payload Account, reply *DataOuputs) error {

	//active or not account
	if initializeAccount(b, reply) != true {
		return nil
	}

	// set reply value
	reply.Account.ActiveCard = payload.ActiveCard
	reply.Account.AvailableLimit = payload.AvailableLimit

	// set struct DB
	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	b.database["lastime"] = ""

	fmt.Printf("Birds : %+v", reply)
	return nil

}

func (b *Ouputs) AddTransaction(payload Transaction, reply *DataOuputs) error {

	// set reply value
	reply.Account.ActiveCard = true
	reply.Violations = []string{""}

	// approve the transaction amount
	aprovalAmount(reply, payload.Amount)

	// Check time
	checkTime(reply, payload.Time, b.database["lastime"].(string))

	// set struct DB
	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	b.database["lastime"] = payload.Time

	return nil

}

//active or not account
func initializeAccount(b *Ouputs, reply *DataOuputs) bool {

	if b.database["account"] == nil {
		reply.Violations = append(reply.Violations, "account-not-initialized")
		return false

	}
	dataAccount := b.database["account"].(Account)
	reply.Account.AvailableLimit = dataAccount.AvailableLimit
	if dataAccount.ActiveCard != true {
		reply.Violations = append(reply.Violations, "account-not-initialized")
		return false

	}
	reply.Violations = []string{""}

	return true

}

//colse Account
func (b *Ouputs) Ending(payload Transaction, reply *DataOuputs) error {
	b.database = make(map[string]interface{})

	return nil

}

// approve the transaction amount
func aprovalAmount(reply *DataOuputs, amaunt float64) float64 {

	newAmount := reply.Account.AvailableLimit - amaunt

	if newAmount < 0 {
		reply.Violations = append(reply.Violations, "insufficient-limit")
	} else {
		reply.Account.AvailableLimit = newAmount
	}

	return newAmount
}

//check the times
func checkTime(reply *DataOuputs, newTime string, lastTime string) {

	if lastTime == "" {
		lastTime = "2006-01-02T15:04:05.000Z"
	}

	const layout = "2006-01-02T15:04:05.000000Z"

	tt := strings.Split(lastTime, "Z")
	lastTime = tt[0] + "000Z"

	tt = strings.Split(newTime, "Z")
	newTime = tt[0] + "000Z"

	tlast, _ := time.Parse(layout, lastTime)
	tnew, _ := time.Parse(layout, newTime)

	fmt.Println(tlast.Unix() - tnew.Unix())

	if (tnew.Unix() - tlast.Unix()) < 10 {
		reply.Violations = append(reply.Violations, "high-frequency-small-interval")
	}

}

// NewCard function returns a new instance of card (pointer).
func NewCard() *Ouputs {
	return &Ouputs{
		database: make(map[string]interface{}),
	}
}
