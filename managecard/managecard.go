package managecard

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Account struct represents a card
type Account struct {
	ActiveCard     bool
	AvailableLimit float64
}

//Transaction struct represents a Transaction
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

// Add methods adds an account to the struct DB (procedure).
func (b *Ouputs) AddAccount(payload Account, reply *DataOuputs) error {

	// set reply value
	reply.Account.ActiveCard = payload.ActiveCard
	reply.Account.AvailableLimit = payload.AvailableLimit

	// set struct DB
	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	b.database["lastime1"] = "2006-01-02T15:04:05.000Z"
	b.database["lastime2"] = "2006-01-02T15:04:05.000Z"
	b.database["lastime3"] = "2006-01-02T15:04:05.000Z"
	b.database["merchant"] = make(map[string]string)
	b.database["amount"] = float64(0)
	b.database["nt"] = int(0)

	fmt.Printf("Account : %+v \n", reply.Violations)
	return nil

}

// Analyze the transaction (procedure).
func (b *Ouputs) AddTransaction(payload Transaction, reply *DataOuputs) error {

	//active or not account
	if initializeAccount(b, reply) != true {
		fmt.Printf("Account : %+v \n", reply.Violations)
		return nil
	}

	nt := b.database["nt"].(int)
	if nt < 3 {
		nt++
	}
	b.database["nt"] = nt

	// set reply value
	reply.Account.ActiveCard = true
	reply.Violations = []string{""}

	if nt == 1 {
		b.database["lastime1"] = payload.Time
	} else if nt == 2 {
		b.database["lastime2"] = payload.Time
	} else if nt == 3 {
		b.database["lastime3"] = payload.Time
	}

	// Check time
	if aprovalMerchan(reply, b.database, payload) == true && checkTime(reply, b.database) == true {

		// approve the transaction amount
		aprovalAmount(reply, payload.Amount)
	}

	if nt == 3 {
		b.database["lastime1"] = b.database["lastime2"]
		b.database["lastime2"] = b.database["lastime3"]
		b.database["lastime3"] = "2006-01-02T15:04:05.000Z"
	}

	// set struct DB
	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	merchants := b.database["merchant"].(map[string]string)
	s := fmt.Sprintf("%f", payload.Amount)
	merchants[payload.Merchant] = payload.Time + ";" + s
	b.database["merchant"] = merchants
	b.database["amount"] = payload.Amount

	fmt.Printf("Transaction : %+v \n", reply.Violations)

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
		reply.Violations = append(reply.Violations, "card-not-active")
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

	return reply.Account.AvailableLimit
}

// approve the transaction amount
func aprovalMerchan(reply *DataOuputs, database map[string]interface{}, payload Transaction) bool {

	merchants := database["merchant"].(map[string]string)

	if merchants[payload.Merchant] == "" {
		return true
	}

	datam := strings.Split(merchants[payload.Merchant], ";")
	difTime := timesString(payload.Time, datam[0])
	amount, _ := strconv.ParseFloat(datam[1], 32)
	if difTime <= 120 && payload.Amount == amount {
		reply.Violations = append(reply.Violations, "doubled-transaction")
		return false

	}

	return true

}

func timesString(newTime string, lastTime string) int64 {

	const layout = "2006-01-02T15:04:05.000000Z"

	tt := strings.Split(lastTime, "Z")
	lastTime = tt[0] + "000Z"

	tt = strings.Split(newTime, "Z")
	newTime = tt[0] + "000Z"

	tlast, _ := time.Parse(layout, lastTime)
	tnew, _ := time.Parse(layout, newTime)

	return tnew.Unix() - tlast.Unix()

}

//check the time
func checkTime(reply *DataOuputs, database map[string]interface{}) bool {

	if database["nt"].(int) < 3 {
		return true
	}

	difTime := timesString(database["lastime3"].(string), database["lastime1"].(string))

	if difTime <= 120 {
		reply.Violations = append(reply.Violations, "high-frequency-small-interval")
		return false
	}

	return true

}

// NewCard function returns a new instance of card (pointer).
func NewCard() *Ouputs {
	return &Ouputs{
		database: make(map[string]interface{}),
	}
}
