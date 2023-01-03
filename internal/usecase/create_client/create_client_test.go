package create_client

import (
	"testing"

	"github.com/brunomann/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	mockRepository := &ClientGatewayMock{}
	mockRepository.On("Save", mock.Anything).Return(nil)
	useCase := NewCreateClientUseCase(mockRepository)

	output, err := useCase.Execute(CreateClientInputDTO{
		Name:  "Bruno Mann",
		Email: "bruno@bruno.com",
	})

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.NotEmpty(t, output.CreatedAt)
	assert.Equal(t, output.Name, "Bruno Mann")
	assert.Equal(t, output.Email, "bruno@bruno.com")
	mockRepository.AssertExpectations(t)
	mockRepository.AssertNumberOfCalls(t, "Save", 1)
}
