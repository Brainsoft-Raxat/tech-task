package repository

import (
	"context"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/app/connection"
	"github.com/Brainsoft-Raxat/tech-task/internal/models"

	"go.uber.org/zap"
)

type AccountRepository interface {
	CreateAccount(ctx context.Context, account models.Account) (models.Account, error)
	GetAllAccounts(ctx context.Context) ([]models.Account, error)
	GetAccountByID(ctx context.Context, id string) (models.Account, error)
	UpdateAccountByID(ctx context.Context, id string, account models.Account) (models.Account, error)
	DeleteAccountByID(ctx context.Context, id string) error
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, transaction models.Transaction) (models.Transaction, error)
	GetAllTransactionsByAccountID(ctx context.Context, accountID string) ([]models.Transaction, error)
	GetTransactionByID(ctx context.Context, id string) (models.Transaction, error)
	UpdateTransactionByID(ctx context.Context, id string, transaction models.Transaction) (models.Transaction, error)
	DeleteTransactionByID(ctx context.Context, id string) error
}

type Repository struct {
	AccountRepository
	TransactionRepository
}

func New(conn *connection.Connection, cfg *config.Configs, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		AccountRepository:     NewAccountRepository(conn.Postgres, cfg, logger),
		TransactionRepository: NewTransactionRepository(conn.Postgres, cfg, logger),
	}
}
