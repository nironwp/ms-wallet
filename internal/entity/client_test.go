package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "john@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "john@email.com", client.Email)
	assert.NotEmpty(t, client.ID)
	assert.NotZero(t, client.CreatedAt)
	assert.NotZero(t, client.UpdatedAt)
}

func TestCreatenewClientWhenArgsAreInvalid(t *testing.T) {
	_, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Equal(t, "name is required", err.Error())

	_, err = NewClient("John", "")
	assert.NotNil(t, err)
	assert.Equal(t, "email is required", err.Error())
}

func TestUpdateAClientWithValidParameters(t *testing.T) {
	client, err := NewClient("John", "john@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)

	err = client.Update("Pedro", "")

	assert.Nil(t, err)
	assert.Equal(t, client.Name, "Pedro")

	err = client.Update("", "pedro@gmail.com")

	assert.Nil(t, err)
	assert.Equal(t, client.Email, "pedro@gmail.com")
}

func TestAddAAccountToClient(t *testing.T) {
	client, err := NewClient("John", "john@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)

	account, err := NewAccount(client)

	err = client.AddAccount(account)

	assert.Nil(t, err)
}

func TestAddAccountShouldThrowAnErrorWhenClientIsNotTheOwnerOfTheAccount(t *testing.T) {
	client, err := NewClient("John", "john@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)

	client2, err := NewClient("Pedro", "pedro@email.com")
	assert.Nil(t, err)
	assert.NotNil(t, client2)

	account, err := NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)

	err = client2.AddAccount(account)

	assert.Equal(t, err.Error(), "This client is not owner of this account")
}
