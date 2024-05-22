package service

import (
	"context"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/data"
	"github.com/Brainsoft-Raxat/tech-task/internal/repository"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AccountService interface {
	CreateAccount(ctx context.Context, req data.CreateAccountRequest) (data.CreateAccountResponse, error)
	GetAllAccounts(ctx context.Context, req data.GetAllAccountsRequest) (resp data.GetAllAccountsResponse, err error)
	GetAccountByID(ctx context.Context, req data.GetAccountByIDRequest) (resp data.GetAccountByIDResponse, err error)
	UpdateAccount(ctx context.Context, req data.UpdateAccountRequest) (resp data.UpdateAccountResponse, err error)
	DeleteAccount(ctx context.Context, req data.DeleteAccountRequest) (resp data.DeleteAccountResponse, err error)
}

type TransactionService interface {
	CreateTransaction(ctx context.Context, req data.CreateTransactionRequest) (resp data.CreateTransactionResponse, err error)
	GetAllTransactionsByAccountID(ctx context.Context, req data.GetAllTransactionsByAccountIDRequest) (resp data.GetAllTransactionsByAccountIDResponse, err error)
	GetTransactionByID(ctx context.Context, req data.GetTransactionByIDRequest) (resp data.GetTransactionByIDResponse, err error)
	DeleteTransaction(ctx context.Context, req data.DeleteTransactionRequest) (resp data.DeleteTransactionResponse, err error)
}

type Service struct {
	AccountService
	TransactionService
}

func New(repos *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	validator := validator.New()

	srv := &Service{
		AccountService:     NewAccountService(repos, cfg, logger, validator),
		TransactionService: NewTransactionService(repos, cfg, logger, validator),
	}

	return srv
}
