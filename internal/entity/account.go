package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ac *Account) Validate() error {
	if ac.Client == nil {
		return errors.New("account need a client")
	}

	return nil
}

func (ac *Account) Credit(amount float64) {
	ac.Balance += amount
	ac.UpdatedAt = time.Now()
}

func (ac *Account) Debit(amount float64) {
	ac.Balance -= amount
	ac.UpdatedAt = time.Now()
}

func NewAccount(client *Client) (*Account, error) {
	account := Account{
		ID:        uuid.NewString(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := account.Validate()

	if err != nil {
		return nil, err
	}

	return &account, nil
}
