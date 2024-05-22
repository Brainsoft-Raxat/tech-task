package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/tech-task/internal/data"

	"github.com/labstack/echo/v4"
)

// CreateAccount godoc
// @Summary Create account
// @Description Create account
// @Tags account
// @Accept json
// @Produce json
// @Param request body data.CreateAccountRequest true "Create account"
// @Success 200 {object} data.CreateAccountResponse
// @Router /account [post]
func (h *handler) CreateAccount(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.CreateAccountRequest
	if err := c.Bind(&req); err != nil {
		return HandleEcho(c, err)
	}

	resp, err := h.service.AccountService.CreateAccount(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// GetAllAccounts godoc
// @Summary Get all accounts
// @Description Get all accounts
// @Tags account
// @Produce json
// @Success 200 {object} data.GetAllAccountsResponse
// @Router /account [get]
func (h *handler) GetAllAccounts(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.GetAllAccountsRequest

	resp, err := h.service.AccountService.GetAllAccounts(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// GetAccountByID godoc
// @Summary Get account by ID
// @Description Get account by ID
// @Tags account
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} data.GetAccountByIDResponse
// @Router /account/{id} [get]
func (h *handler) GetAccountByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.GetAccountByIDRequest

	req.ID = c.Param("id")

	resp, err := h.service.AccountService.GetAccountByID(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// UpdateAccount godoc
// @Summary Update account
// @Description Update account
// @Tags account
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param request body data.UpdateAccountRequest true "Update account"
// @Success 200 {object} data.UpdateAccountResponse
// @Router /account/{id} [put]
func (h *handler) UpdateAccount(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.UpdateAccountRequest
	if err := c.Bind(&req); err != nil {
		return HandleEcho(c, err)
	}

	req.ID = c.Param("id")

	resp, err := h.service.AccountService.UpdateAccount(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteAccount godoc
// @Summary Delete account
// @Description Delete account
// @Tags account
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} data.DeleteAccountResponse
// @Router /account/{id} [delete]
func (h *handler) DeleteAccount(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.DeleteAccountRequest

	req.ID = c.Param("id")

	resp, err := h.service.AccountService.DeleteAccount(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}
