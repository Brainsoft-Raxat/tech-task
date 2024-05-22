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
	"go.uber.org/zap"
)

type accountService struct {
	cfg         *config.Configs
	logger      *zap.SugaredLogger
	validator   *validator.Validate
	accountRepo repository.AccountRepository
}

func NewAccountService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger, validator *validator.Validate) AccountService {
	return &accountService{
		cfg:         cfg,
		logger:      logger,
		validator:   validator,
		accountRepo: repo.AccountRepository,
	}
}

func (s *accountService) CreateAccount(ctx context.Context, req data.CreateAccountRequest) (resp data.CreateAccountResponse, err error) {
	s.logger.Infow("CreateAccount", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("CreateAccount", "err", err)
			return
		}
		s.logger.Infow("CreateAccount", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	account := models.Account{
		Name:    req.Name,
		Balance: req.Balance,
	}

	account, err = s.accountRepo.CreateAccount(ctx, account)
	if err != nil {
		return
	}

	resp = data.CreateAccountResponse{
		Account: account,
	}

	return
}

func (s *accountService) GetAllAccounts(ctx context.Context, req data.GetAllAccountsRequest) (resp data.GetAllAccountsResponse, err error) {
	s.logger.Infow("GetAllAccounts", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("GetAllAccounts", "err", err)
			return
		}
		s.logger.Infow("GetAllAccounts", "response", resp)
	}()

	accounts, err := s.accountRepo.GetAllAccounts(ctx)
	if err != nil {
		return
	}

	resp = data.GetAllAccountsResponse{
		Accounts: accounts,
	}

	return
}

func (s *accountService) GetAccountByID(ctx context.Context, req data.GetAccountByIDRequest) (resp data.GetAccountByIDResponse, err error) {
	s.logger.Infow("GetAccountByID", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("GetAccountByID", "err", err)
			return
		}
		s.logger.Infow("GetAccountByID", "response", resp)
	}()

	account, err := s.accountRepo.GetAccountByID(ctx, req.ID)
	if err != nil {
		return
	}

	resp = data.GetAccountByIDResponse{
		Account: account,
	}

	return
}

func (s *accountService) UpdateAccount(ctx context.Context, req data.UpdateAccountRequest) (resp data.UpdateAccountResponse, err error) {
	s.logger.Infow("UpdateAccount", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("UpdateAccount", "err", err)
			return
		}
		s.logger.Infow("UpdateAccount", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	account := models.Account{
		Name:    req.Name,
		Balance: req.Balance,
	}

	account, err = s.accountRepo.UpdateAccountByID(ctx, req.ID, account)
	if err != nil {
		return
	}

	resp = data.UpdateAccountResponse{
		Account: account,
	}

	return
}

func (s *accountService) DeleteAccount(ctx context.Context, req data.DeleteAccountRequest) (resp data.DeleteAccountResponse, err error) {
	s.logger.Infow("DeleteAccount", "request", req)
	defer func() {
		if err != nil {
			s.logger.Errorw("DeleteAccount", "err", err)
			return
		}
		s.logger.Infow("DeleteAccount", "response", resp)
	}()

	err = s.validator.StructCtx(ctx, req)
	if err != nil {
		err = apperror.NewErrorInfo(ctx, errcodes.InvalidRequest, err.Error()).SetMessage(err.Error())
		return
	}

	err = s.accountRepo.DeleteAccountByID(ctx, req.ID)
	if err != nil {
		return
	}

	return
}
