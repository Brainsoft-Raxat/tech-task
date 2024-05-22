package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/models"
	"github.com/Brainsoft-Raxat/tech-task/pkg/apperror"
	"github.com/Brainsoft-Raxat/tech-task/pkg/errcodes"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type transactionRepository struct {
	client *sqlx.DB
	cfg    *config.Configs
	logger *zap.SugaredLogger
}

func NewTransactionRepository(client *sqlx.DB, cfg *config.Configs, logger *zap.SugaredLogger) TransactionRepository {
	return &transactionRepository{
		client: client,
		cfg:    cfg,
		logger: logger,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, transaction models.Transaction) (models.Transaction, error) {
	tx, err := r.client.BeginTxx(ctx, nil)
	if err != nil {
		return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("failed to begin Tx: %v", err))
	}

	defer func() {
		if p := recover(); p != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				r.logger.Errorf("rollback error: %v", rollbackErr)
			}
			panic(p)
		} else if err != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				r.logger.Errorf("rollback error: %v", rollbackErr)
			}
		} else {
			err = tx.Commit()
		}
	}()

	var account models.Account
	err = tx.GetContext(ctx, &account, "SELECT id, name, balance, created_at, updated_at FROM accounts WHERE id = $1", transaction.AccountID)
	if err != nil {
		return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("failed to get account: %v", err))
	}

	query := `
		INSERT INTO transactions (value, account_id, group_type, created_at, updated_at)
		VALUES (:value, :account_id, :group_type, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id, value, account_id, group_type, created_at, updated_at;
	`
	namedArgs := map[string]interface{}{
		"value":       transaction.Value,
		"account_id":  transaction.AccountID,
		"group_type":  transaction.GroupType,
	}

	switch transaction.GroupType {
	case models.GroupTypeTransfer:
		account.Balance -= transaction.Value
		if account.Balance < 0 {
			return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "insufficient funds").SetMessage("insufficient funds")
		}

		var account2 models.Account
		err = tx.GetContext(ctx, &account2, "SELECT id, name, balance, created_at, updated_at FROM accounts WHERE id = $1", transaction.Account2ID)
		if err != nil {
			return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("failed to get account2: %v", err))
		}

		account2.Balance += transaction.Value
		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", account2.Balance, account2.ID)
		if err != nil {
			return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("failed to update account2 balance: %v", err))
		}

		query = `
			INSERT INTO transactions (value, account_id, group_type, account2_id, created_at, updated_at)
			VALUES (:value, :account_id, :group_type, :account2_id, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
			RETURNING id, value, account_id, group_type, account2_id, created_at, updated_at;
		`
		namedArgs = map[string]interface{}{
			"value":       transaction.Value,
			"account_id":  transaction.AccountID,
			"group_type":  transaction.GroupType,
			"account2_id": transaction.Account2ID,
		}

	case models.GroupTypeIncome:
		transaction.Account2ID = uuid.Nil
		account.Balance += transaction.Value

	case models.GroupTypeOutcome:
		transaction.Account2ID = uuid.Nil
		account.Balance -= transaction.Value
		if account.Balance < 0 {
			return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "insufficient funds").SetMessage("insufficient funds")
		}

	default:
		return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "invalid group type")
	}

	_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = $1 WHERE id = $2", account.Balance, account.ID)
	if err != nil {
		return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("failed to update account balance: %v", err))
	}

	rows, err := tx.NamedQuery(query, namedArgs)
	if err != nil {
		r.logger.Errorf("failed to create transaction: %v", err)
		return models.Transaction{}, err
	}
	defer rows.Close()

	var newTransaction models.Transaction
	if rows.Next() {
		err = rows.StructScan(&newTransaction)
		if err != nil {
			r.logger.Errorf("failed to scan transaction: %v", err)
			return models.Transaction{}, err
		}
	} else {
		return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "no rows returned")
	}

	return newTransaction, nil
}

func (r *transactionRepository) GetAllTransactionsByAccountID(ctx context.Context, accountID string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := `
		SELECT id, value, account_id, group_type, account2_id, created_at, updated_at
		FROM transactions
		WHERE account_id = $1 OR account2_id = $1
	`
	rows, err := r.client.QueryxContext(ctx, query, accountID)
	if err != nil {
		r.logger.Errorf("failed to get transactions by account id: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.StructScan(&transaction)
		if err != nil {
			r.logger.Errorf("failed to scan transaction: %v", err)
			return nil, err
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactionByID(ctx context.Context, id string) (models.Transaction, error) {
	var transaction models.Transaction

	query := `
		SELECT id, value, account_id, group_type, account2_id, created_at, updated_at
		FROM transactions
		WHERE id = $1
	`
	err := r.client.QueryRowxContext(ctx, query, id).StructScan(&transaction)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Transaction{}, apperror.NewErrorInfo(ctx, errcodes.NotFoundError, err.Error()).SetMessage("transaction not found")
		}
		return models.Transaction{}, err
	}

	return transaction, nil
}

func (r *transactionRepository) UpdateTransactionByID(ctx context.Context, id string, transaction models.Transaction) (models.Transaction, error) {
	query := `
		UPDATE transactions
		SET value = :value, updated_at = CURRENT_TIMESTAMP
		WHERE id = :id
		RETURNING id, created_at, updated_at;
	`
	namedArgs := map[string]interface{}{
		"id":    id,
		"value": transaction.Value,
	}

	var updatedTransaction models.Transaction
	err := r.client.QueryRowxContext(ctx, query, namedArgs).StructScan(&updatedTransaction)
	if err != nil {
		r.logger.Errorf("failed to update transaction by id: %v", err)
		return models.Transaction{}, err
	}

	return updatedTransaction, nil
}

func (r *transactionRepository) DeleteTransactionByID(ctx context.Context, id string) error {
	query := `
		DELETE FROM transactions
		WHERE id = $1
	`
	_, err := r.client.ExecContext(ctx, query, id)
	if err != nil {
		r.logger.Errorf("failed to delete transaction by id: %v", err)
		return err
	}

	return nil
}
