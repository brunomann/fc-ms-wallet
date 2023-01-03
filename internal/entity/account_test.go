package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := newClient("Bruno Ramos", "bruno@gmail.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestDepositInAccount(t *testing.T) {
	client, _ := newClient("Bruno Ramos", "bruno@gmail.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, account.Balance, 100.0)
}

func TestDebitInAccount(t *testing.T) {
	client, _ := newClient("Bruno Ramos", "bruno@gmail.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, account.Balance, 50.0)
}
