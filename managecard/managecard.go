package managecard

import (
	"fmt"
	"time"
)

// Student struct represents a student.
type Student struct {
	ID                  int
	FirstName, LastName string
}

//Cards struct represents a card
type Cards struct {
	ActiveCard     bool
	AvailableLimit float64
}

type Violations struct {
	Violations []string
}

type Account struct {
	ActiveCard     bool
	AvailableLimit float64
}

type ReplyCards struct {
	Account Account
}

type OutputTransaction struct {
	Account    Account
	Violations Violations
}

type Transaction struct {
	Merchant string
	Amount   float64
	Time     string
}

// FullName returns the fullname of the student.
func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

// College struct represents a college.
type Ouputs struct {
	database map[string]interface{} // private

}

// College struct represents a college.
type DataOuputs struct {
	Account    Account
	Violations []string
}

/*---------------*/

func (b *Ouputs) AddAccount(payload Cards, reply *DataOuputs) error {

	// set reply value
	reply.Account.ActiveCard = payload.ActiveCard
	reply.Account.AvailableLimit = payload.AvailableLimit
	reply.Violations = []string{""}

	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	b.database["lastime"] = nil

	fmt.Printf("Birds : %+v", reply)
	return nil

}

func (b *Ouputs) AddTransaction(payload Transaction, reply *DataOuputs) error {

	dataAccount := b.database["account"].(Account)
	reply.Account.AvailableLimit = dataAccount.AvailableLimit

	if dataAccount.ActiveCard != true {
		reply.Violations = append(reply.Violations, "account-not-initialized")
		return nil

	}

	// set reply value
	reply.Account.ActiveCard = true
	reply.Violations = []string{""}

	aprovalAmount(reply, payload.Amount)
	if b.database["lastime"] == nil {
		b.database["lastime"] = payload.Time
	}

	checkTime(reply, b.database["lastime"].(string), payload.Time)

	b.database["account"] = reply.Account
	b.database["violations"] = reply.Violations
	b.database["lastime"] = payload.Time

	return nil

}

func aprovalAmount(reply *DataOuputs, amaunt float64) float64 {

	newAmount := reply.Account.AvailableLimit - amaunt

	if newAmount < 0 {
		reply.Violations = []string{"high-frequency-small-interval"}
	} else {
		reply.Account.AvailableLimit = newAmount
	}

	return newAmount
}

func checkTime(reply *DataOuputs, newTime string, lastTime string) {

	tlast, _ := time.Parse(time.RFC3339, lastTime)
	tnew, _ := time.Parse(time.RFC3339, newTime)

	if (tnew.Unix() - tlast.Unix()) < 10 {
		reply.Violations = []string{"high-frequency-small-interval"}
	}

}

func (c *College) Gge(payload Student, reply *Student) error {
	// check if student already exists in the database
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("student with id '%d' already exists", payload.ID)
	}

	// add student `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Add methods adds a student to the college (procedure).
func (c *College) Add(payload Student, reply *Student) error {

	// check if student already exists in the database
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("student with id '%d' already exists", payload.ID)
	}

	// add student `p` in the database
	c.database[payload.ID] = payload

	// set reply value
	*reply = payload

	// return `nil` error
	return nil
}

// Get methods returns a student with specific id (procedure).
func (c *College) Get(payload int, reply *Student) error {

	// get student with id `p` from the database
	result, ok := c.database[payload]

	// check if student exists in the database
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", payload)
	}

	// set reply value
	*reply = result

	// return `nil` error
	return nil
}

// NewCollege function returns a new instance of College (pointer).
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}

// NewCollege function returns a new instance of College (pointer).
func NewCart() *Ouputs {
	return &Ouputs{
		database: make(map[string]interface{}),
	}
}

// College struct represents a college.
type College struct {
	database map[int]Student // private
}
