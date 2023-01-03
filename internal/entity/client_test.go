package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Bruno Mann", "b@b.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Bruno Mann", client.Name)
	assert.Equal(t, "b@b.com", client.Email)
}

func TestCreateInvalidClient(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Bruno Mann", "b@b.com")
	err := client.Update("Bruno Mann Update", "b2@b.com")
	assert.Nil(t, err)
	assert.Equal(t, "Bruno Mann Update", client.Name)
	assert.Equal(t, "b2@b.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("Bruno Mann", "b@b.com")
	err := client.Update("", "b2@b.com")
	assert.Error(t, err, "Name is required")
	assert.Error(t, err, "Email is required")
}

func TestAddClientToAccount(t *testing.T) {
	client, _ := NewClient("Bruno Mann", "b@b.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddInvalidClientToAccount(t *testing.T) {
	client, _ := NewClient("Bruno Mann", "b@b.com")
	client2, _ := NewClient("Bruno Mann Ramos", "b@b.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	err := client.AddAccount(account)
	assert.Nil(t, err)

	err = client2.AddAccount(account2)
	assert.Nil(t, err)

	err = client.AddAccount(account2)
	assert.Error(t, err, "Account does not belong to client")
}
