package createclient

import (
	"testing"

	"github.com/nironwp/ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (n *ClientGatewayMock) Save(client *entity.Client) error {
	args := n.Called(client)

	return args.Error(0)
}

func (n *ClientGatewayMock) ListAll() ([]*entity.Client, error) {
	args := n.Called()

	return args.Get(0).([]*entity.Client), args.Error(1)
}

func (n *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := n.Called(id)

	return args.Get(0).(*entity.Client), args.Error(1)
}

func TestCreateAClient(t *testing.T) {
	n := &ClientGatewayMock{}
	n.On("Save", mock.Anything).Return(nil)
	input := CreateClientInputDTO{
		Name:  "John",
		Email: "john@gmail.com",
	}

	usecase := NewCreateClientUseCase(n)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	assert.Equal(t, output.Name, input.Name)
	assert.Equal(t, output.Email, input.Email)
	assert.NotZero(t, output.CreatedAt)
	assert.NotZero(t, output.UpdatedAt)
}
