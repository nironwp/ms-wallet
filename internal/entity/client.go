package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	Accounts  []*Account
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}

func (c *Client) Update(name string, email string) error {
	if name != "" {
		c.Name = name
	}
	if email != "" {
		c.Email = email
	}

	err := c.Validate()

	if err != nil {
		return err
	}

	c.UpdatedAt = time.Now()

	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if c.ID != account.Client.ID {
		return errors.New("This client is not owner of this account")
	}

	c.Accounts = append(c.Accounts, account)

	return nil
}

func NewClient(name string, email string) (*Client, error) {

	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := client.Validate()

	if err != nil {
		return nil, err
	}

	return client, nil
}
