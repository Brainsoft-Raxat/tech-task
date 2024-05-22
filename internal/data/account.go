package data

import "github.com/Brainsoft-Raxat/tech-task/internal/models"

type CreateAccountRequest struct {
	Name    string  `json:"name" validate:"required,min=3,max=100"`
	Balance float64 `json:"balance" validate:"required,gt=0"`
}

type CreateAccountResponse struct {
	Account models.Account `json:"account"`
}

type GetAllAccountsRequest struct{}

type GetAllAccountsResponse struct {
	Accounts []models.Account `json:"accounts"`
}

type GetAccountByIDRequest struct {
	ID string `json:"id" validate:"required,uuid"`
}

type GetAccountByIDResponse struct {
	Account models.Account `json:"account"`
}

type UpdateAccountRequest struct {
	ID      string  `json:"id" validate:"required,uuid4"`
	Name    string  `json:"name" validate:"required,min=3,max=100"`
	Balance float64 `json:"balance" validate:"required,gt=0"`
}

type UpdateAccountResponse struct {
	Account models.Account `json:"account"`
}

type DeleteAccountRequest struct {
	ID string `json:"id" validate:"required,uuid4"`
}

type DeleteAccountResponse struct {
	// Account models.Account `json:"account"`
}
