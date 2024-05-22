package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/tech-task/internal/data"

	"github.com/labstack/echo/v4"
)

// CreateTransaction godoc
// @Summary Create transaction
// @Description Create transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Param request body data.CreateTransactionRequest true "Create transaction"
// @Success 200 {object} data.CreateTransactionResponse
// @Router /transaction [post]
func (h *handler) CreateTransaction(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.CreateTransactionRequest
	if err := c.Bind(&req); err != nil {
		return HandleEcho(c, err)
	}

	resp, err := h.service.TransactionService.CreateTransaction(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// GetAllTransactionsByAccountID godoc
// @Summary Get all transactions by account ID
// @Description Get all transactions by account ID
// @Tags transaction
// @Produce json
// @Param account_id path string true "Account ID"
// @Success 200 {object} data.GetAllTransactionsByAccountIDResponse
// @Router /transaction/account/{id} [get]
func (h *handler) GetAllTransactionsByAccountID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.GetAllTransactionsByAccountIDRequest

	req.AccountID = c.Param("id")

	resp, err := h.service.TransactionService.GetAllTransactionsByAccountID(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Description Get transaction by ID
// @Tags transaction
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} data.GetTransactionByIDResponse
// @Router /transaction/{id} [get]
func (h *handler) GetTransactionByID(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.GetTransactionByIDRequest

	req.ID = c.Param("id")

	resp, err := h.service.TransactionService.GetTransactionByID(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteTransaction godoc
// @Summary Delete transaction
// @Description Delete transaction
// @Tags transaction
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} data.DeleteTransactionResponse
// @Router /transaction/{id} [delete]
func (h *handler) DeleteTransaction(c echo.Context) error {
	ctx, cancel := h.context(c)
	defer cancel()

	var req data.DeleteTransactionRequest

	req.ID = c.Param("id")

	resp, err := h.service.TransactionService.DeleteTransaction(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}