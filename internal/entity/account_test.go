package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, err := NewClient("John Doe", "john")

	assert.Nil(t, err)

	account, err := NewAccount(client)
	assert.Equal(t, account.Client.ID, client.ID)
	assert.Equal(t, account.Balance, 0.0)
}

func TestWhenCreateAnAccountWithNotAClientShouldThrowAnError(t *testing.T) {
	account, err := NewAccount(nil)

	assert.Nil(t, account)
	assert.Equal(t, err.Error(), "account need a client")
}

func TestCreditAnAccount(t *testing.T) {
	client, _ := NewClient("John", "john@gmail.com")

	account, err := NewAccount(client)

	assert.Nil(t, err)

	account.Credit(10)

	assert.Equal(t, account.Balance, 10.0)
}

func TestDebitAnAccount(t *testing.T) {
	client, _ := NewClient("John", "john@gmail.com")

	account, err := NewAccount(client)
	assert.Nil(t, err)

	account.Debit(10.0)
	assert.Equal(t, account.Balance, -10.0)

}
