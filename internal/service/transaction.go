package service

import (
	"context"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/data"
	"github.com/Brainsoft-Raxat/tech-task/internal/models"
	"github.com/Brainsoft-Raxat/tech-task/internal/repository"
	"github.com/Brainsoft-Raxat/tech-task/pkg/apperror"
	"github.com/Brainsoft-Raxat/tech-task/pkg/errcodes"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type transactionService struct {
	cfg             *config.Configs
	logger          *zap.SugaredLogger
	validator       *validator.Validate
	transactionRepo repository.TransactionRepository
	accountRepo     repository.AccountRepository
}

func NewTransactionService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger, validator *validator.Validate) TransactionService {
	return &transactionService{
		cfg:             cfg,
		logger:          logger,
		validator:       validator,
		transactionRepo: repo.TransactionRepository,
		accountRepo:     repo.AccountRepository,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req data.CreateTransactionRequest) (resp data.CreateTransactionResponse, err error) {
	s.logger.Infow("CreateTransaction", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("CreateTransaction", "err", err)
			return
		}
		s.logger.Infow("CreateTransaction", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	accountID, err := uuid.Parse(req.AccountID)
	if err != nil {
		return resp, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "invalid account id")
	}

	account2ID := uuid.Nil
	if req.Account2ID != "" {
		account2ID, err = uuid.Parse(req.Account2ID)
		if err != nil {
			return resp, apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, "invalid account2 id")
		}
	}

	transaction := models.Transaction{
		Value:      req.Value,
		AccountID:  accountID,
		GroupType:  req.GroupType,
		Account2ID: account2ID,
	}

	transaction, err = s.transactionRepo.CreateTransaction(ctx, transaction)
	if err != nil {
		return
	}

	resp = data.CreateTransactionResponse{
		Transaction: transaction,
	}

	return
}

func (s *transactionService) GetAllTransactionsByAccountID(ctx context.Context, req data.GetAllTransactionsByAccountIDRequest) (resp data.GetAllTransactionsByAccountIDResponse, err error) {
	s.logger.Infow("GetAllTransactionsByAccountID", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("GetAllTransactionsByAccountID", "err", err)
			return
		}
		s.logger.Infow("GetAllTransactionsByAccountID", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	transactions, err := s.transactionRepo.GetAllTransactionsByAccountID(ctx, req.AccountID)
	if err != nil {
		return
	}

	resp = data.GetAllTransactionsByAccountIDResponse{
		Transactions: transactions,
	}

	return
}

func (s *transactionService) GetTransactionByID(ctx context.Context, req data.GetTransactionByIDRequest) (resp data.GetTransactionByIDResponse, err error) {
	s.logger.Infow("GetTransactionByID", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("GetTransactionByID", "err", err)
			return
		}
		s.logger.Infow("GetTransactionByID", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	transaction, err := s.transactionRepo.GetTransactionByID(ctx, req.ID)
	if err != nil {
		return
	}

	resp = data.GetTransactionByIDResponse{
		Transaction: transaction,
	}

	return
}

// func (s *transactionService) UpdateTransaction(ctx context.Context, req data.UpdateTransactionRequest) (resp data.UpdateTransactionResponse, err error) {
// 	s.logger.Infow("UpdateTransaction", "request", req)
// 	defer func() {
// 		if err != nil {
// 			s.logger.Errorw("UpdateTransaction", "err", err)
// 			return
// 		}
// 		s.logger.Infow("UpdateTransaction", "response", resp)
// 	}()

// 	transaction := models.Transaction{
// 		Value:      req.Value,
// 		AccountID:  req.AccountID,
// 		GroupType:  req.GroupType,
// 		Account2ID: req.Account2ID,
// 	}

// 	transaction, err = s.transactionRepo.UpdateTransactionByID(ctx, req.ID, transaction)
// 	if err != nil {
// 		return
// 	}

// 	resp = data.UpdateTransactionResponse{
// 		Transaction: transaction,
// 	}

// 	return
// }

func (s *transactionService) DeleteTransaction(ctx context.Context, req data.DeleteTransactionRequest) (resp data.DeleteTransactionResponse, err error) {
	s.logger.Infow("DeleteTransaction", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("DeleteTransaction", "err", err)
			return
		}
		s.logger.Infow("DeleteTransaction", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	err = s.transactionRepo.DeleteTransactionByID(ctx, req.ID)
	if err != nil {
		return
	}

	resp = data.DeleteTransactionResponse{}

	return
}
