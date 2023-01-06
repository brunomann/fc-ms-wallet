package create_transaction

import (
	"testing"

	"github.com/brunomann/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Bruno Mann", "bruno@bruno.com")
	account := entity.NewAccount(client)
	account.Credit(1000)

	client2, _ := entity.NewClient("Bruno Mann2", "bruno2@bruno.com")
	account2 := entity.NewAccount(client2)
	account2.Credit(1000)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindById", account.ID).Return(account, nil)
	mockAccount.On("FindById", account2.ID).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account.ID,
		AccountIDTo:   account2.ID,
		Amount:        500,
	}

	transactionUseCase := NewCreateTransactionUseCase(mockTransaction, mockAccount)

	output, err := transactionUseCase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	mockAccount.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertExpectations(t)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
