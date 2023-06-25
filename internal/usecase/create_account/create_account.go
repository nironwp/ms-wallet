package createaccount

import (
	"github.com/nironwp/ms-wallet/internal/entity"
	"github.com/nironwp/ms-wallet/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccontUseCase struct {
	AccountCateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccontUseCase {
	return &CreateAccontUseCase{
		AccountCateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (usecase *CreateAccontUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := usecase.ClientGateway.Get(input.ClientID)

	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(client)

	if err != nil {
		return nil, err
	}

	err = usecase.AccountCateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
