package createaccount

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

type AccountGatewayMock struct {
	mock.Mock
}

func (mock *AccountGatewayMock) Save(account *entity.Account) error {
	args := mock.Called(account)

	return args.Error(0)
}
func (mock *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := mock.Called(id)

	return args.Get(0).(*entity.Account), args.Error(1)
}

func (mock *AccountGatewayMock) UpdateBalance(account *entity.Account) error {
	args := mock.Called(account)

	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, err := entity.NewClient("John", "john@gmail.com")

	assert.Nil(t, err)

	assert.NotNil(t, client)

	mockClient := &ClientGatewayMock{}
	mockClient.On("Get", client.ID).Return(client, nil)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("Save", mock.Anything).Return(nil)

	usecase := NewCreateAccountUseCase(mockAccount, mockClient)
	input := CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
}
