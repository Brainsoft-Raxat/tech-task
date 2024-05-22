package repository

import (
	"context"
	"database/sql"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/models"
	"github.com/Brainsoft-Raxat/tech-task/pkg/apperror"
	"github.com/Brainsoft-Raxat/tech-task/pkg/errcodes"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type accountRepository struct {
	client *sqlx.DB
	cfg    *config.Configs
	logger *zap.SugaredLogger
}

func NewAccountRepository(client *sqlx.DB, cfg *config.Configs, logger *zap.SugaredLogger) *accountRepository {
	return &accountRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *accountRepository) CreateAccount(ctx context.Context, account models.Account) (models.Account, error) {
	query := `
		INSERT INTO accounts (name, balance) 
		VALUES (:name, :balance) 
		RETURNING id, name, balance, created_at, updated_at
	`
	namedArgs := map[string]interface{}{
		"name":    account.Name,
		"balance": account.Balance,
	}

	var newAccount models.Account
	rows, err := r.client.NamedQueryContext(ctx, query, namedArgs)
	if err != nil {
		return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&newAccount)
		if err != nil {
			return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
		}
	}

	return newAccount, nil
}

func (r *accountRepository) GetAllAccounts(ctx context.Context) ([]models.Account, error) {
	var accounts []models.Account

	rows, err := r.client.QueryContext(ctx, "SELECT id, name, balance, created_at, updated_at FROM accounts")
	if err != nil {
		return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var account models.Account
		err := rows.Scan(&account.ID, &account.Name, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
		if err != nil {
			return nil, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *accountRepository) GetAccountByID(ctx context.Context, id string) (models.Account, error) {
	var account models.Account

	row := r.client.QueryRowContext(ctx,
		"SELECT id, name, balance, created_at, updated_at FROM accounts WHERE id = $1",
		id,
	)

	err := row.Scan(&account.ID, &account.Name, &account.Balance, &account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, err.Error()).SetMessage("account not found")
		}
		return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}

	return account, nil
}

func (r *accountRepository) UpdateAccountByID(ctx context.Context, id string, account models.Account) (models.Account, error) {
	query := `
		UPDATE accounts 
		SET name = :name, balance = :balance 
		WHERE id = :id 
		RETURNING id, name, balance, created_at, updated_at
	`
	namedArgs := map[string]interface{}{
		"id":      id,
		"name":    account.Name,
		"balance": account.Balance,
	}

	var updatedAccount models.Account
	rows, err := r.client.NamedQueryContext(ctx, query, namedArgs)
	if err != nil {
		r.logger.Errorw("UpdateAccountByID", "err", err)
		return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&updatedAccount)
		if err != nil {
			r.logger.Errorw("UpdateAccountByID", "err", err)
			return models.Account{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
		}
	}

	return updatedAccount, nil
}

func (r *accountRepository) DeleteAccountByID(ctx context.Context, id string) error {
	_, err := r.client.ExecContext(ctx, "DELETE FROM accounts WHERE id = $1", id)
	if err != nil {
		return apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}

	return nil
}
