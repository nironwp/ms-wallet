package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {

	transaction := &Transaction{
		ID:          uuid.NewString(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := transaction.Validate()

	if err != nil {
		return nil, err
	}
	accountFrom.Debit(amount)
	accountTo.Credit(amount)
	return transaction, nil
} 

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must bre greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("Insufficient funds")
	}

	return nil
}
