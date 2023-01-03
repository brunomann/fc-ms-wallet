package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionWhenClientHasBalance(t *testing.T) {
	client, _ := newClient("Bruno Mann", "b@b.com")
	client2, _ := newClient("Bruno Mann Ramos", "b@b.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	err := client.AddAccount(account)
	assert.Nil(t, err)

	err = client2.AddAccount(account2)
	assert.Nil(t, err)

	account.Credit(100)
	account2.Credit(200)

	transaction, err := NewTransaction(account, account2, float64(50))

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account.Balance, float64(50))
	assert.Equal(t, account2.Balance, float64(250))

}

func TestTransactionWhenClientNotHasBalance(t *testing.T) {
	client, _ := newClient("Bruno Mann", "b@b.com")
	client2, _ := newClient("Bruno Mann Ramos", "b@b.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	err := client.AddAccount(account)
	assert.Nil(t, err)

	err = client2.AddAccount(account2)
	assert.Nil(t, err)

	account.Credit(50)
	account2.Credit(200)

	transaction, err := NewTransaction(account, account2, float64(100))

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "Account not have funds to transaction")
	assert.Equal(t, account.Balance, float64(50))
	assert.Equal(t, account2.Balance, float64(200))

}

func TestTransactionWithZeroValue(t *testing.T) {
	client, _ := newClient("Bruno Mann", "b@b.com")
	client2, _ := newClient("Bruno Mann Ramos", "b@b.com")

	account := NewAccount(client)
	account2 := NewAccount(client2)

	err := client.AddAccount(account)
	assert.Nil(t, err)

	err = client2.AddAccount(account2)
	assert.Nil(t, err)

	account.Credit(50)
	account2.Credit(200)

	transaction, err := NewTransaction(account, account2, float64(0))

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "Amount must be greater than zero")
	assert.Equal(t, account.Balance, float64(50))
	assert.Equal(t, account2.Balance, float64(200))

}
