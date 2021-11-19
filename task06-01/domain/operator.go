package domain

import (
	"sync"
)

type Operator struct {
	sync.Mutex
	Username string
	balance  float64
}

func NewOperator(username string) (*Operator, error) {
	// TODO: checking that operator with username doesn't exist yet.
	return &Operator{Username: username, balance: 0}, nil
}

func (o *Operator) Add(amount float64) {
	o.Lock()
	o.balance += amount
	for i := 0; i < 1e8; i += 1 {
	}
	o.Unlock()
}

// Withdraw money to operator card.
// Transfer to operator card is not implemented yet, so we just decrease balance by amount :)
func (o *Operator) Withdraw(amount float64) {
	o.Lock()
	o.balance -= amount
	for i := 0; i < 1e8; i += 1 {
	}
	o.Unlock()
}

func (o *Operator) GetBalance() float64 {
	o.Lock()
	defer o.Unlock()
	return o.balance
}
