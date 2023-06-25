package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateATransaction(t *testing.T) {
	clientFrom, _ := NewClient("John", "john@gmail.com")
	assert.NotNil(t, clientFrom)

	clientTo, _ := NewClient("Pedro", "pedro@gmail.com")
	assert.NotNil(t, clientTo)

	accountFrom, _ := NewAccount(clientFrom)
	assert.NotNil(t, accountFrom)

	accountFrom.Credit(10)

	accountTo, _ := NewAccount(clientTo)
	assert.NotNil(t, accountTo)

	transaction, err := NewTransaction(accountFrom, accountTo, 10)

	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}

func TestShouldThrowAnErrorWhenAccountFromNotHaveBalance(t *testing.T) {
	clientFrom, _ := NewClient("John", "john@gmail.com")
	assert.NotNil(t, clientFrom)

	clientTo, _ := NewClient("Pedro", "pedro@gmail.com")
	assert.NotNil(t, clientTo)

	accountFrom, _ := NewAccount(clientFrom)
	assert.NotNil(t, accountFrom)

	accountTo, _ := NewAccount(clientTo)
	assert.NotNil(t, accountTo)

	transaction, err := NewTransaction(accountFrom, accountTo, 10)

	assert.Nil(t, transaction)
	assert.Equal(t, "Insufficient funds", err.Error())
}

func TestShouldThrowAnErrorWhenInvalidAmount(t *testing.T) {
	clientFrom, _ := NewClient("John", "john@gmail.com")
	assert.NotNil(t, clientFrom)

	clientTo, _ := NewClient("Pedro", "pedro@gmail.com")
	assert.NotNil(t, clientTo)

	accountFrom, _ := NewAccount(clientFrom)
	assert.NotNil(t, accountFrom)

	accountTo, _ := NewAccount(clientTo)
	assert.NotNil(t, accountTo)

	transaction, err := NewTransaction(accountFrom, accountTo, 0)

	assert.Nil(t, transaction)
	assert.Equal(t, "amount must bre greater than zero", err.Error())
}
