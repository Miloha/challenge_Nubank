package managecard

import (
	"fmt"
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
	Violations Violations
}

/*---------------*/

func (b *Ouputs) AddAccount(payload Cards, reply *DataOuputs) error {

	var outs ReplyCards
	outs.Account.ActiveCard = payload.ActiveCard
	outs.Account.AvailableLimit = payload.AvailableLimit
	var vi Violations
	vi.Violations = append(vi.Violations, "siii")

	// set reply value

	b.database["account"] = outs
	b.database["violations"] = vi

	reply.Account.ActiveCard = payload.ActiveCard
	reply.Account.AvailableLimit = payload.AvailableLimit
	reply.Violations.Violations = vi.Violations
	fmt.Printf("Birds : %+v", reply)
	return nil

}

func (b *Ouputs) AddTransaction(payload Transaction, reply *OutputTransaction) error {

	dataAccoun := b.database["account"].(map[string]interface{})

	if dataAccoun["ActiveCard"].(bool) != true {
		reply.Violations.Violations = append(reply.Violations.Violations, "account-not-initialized")
		reply.Account.AvailableLimit = dataAccoun["AvailableLimit"].(float64)
		return nil

	}

	var vi Violations
	reply.Account.ActiveCard = true
	reply.Account.AvailableLimit = dataAccoun["AvailableLimit"].(float64) - payload.Amount
	reply.Violations.Violations = vi.Violations
	b.database["account"] = reply.Account
	b.database["violations"] = vi

	// set reply value

	fmt.Printf("Birds : %+v", reply)
	return nil

}

func (b *Ouputs) Output(payload Ouputs, reply *Ouputs) error {

	reply = b

	return nil

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
