package handler

import (
	"context"
	"net/http"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/service"
	"github.com/Brainsoft-Raxat/tech-task/pkg/apperror"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
)

type handler struct {
	service *service.Service
	cfg     *config.Configs
	logger  *zap.SugaredLogger
}

type Handler interface {
	SetAPI(e *echo.Echo)
}

func New(services *service.Service, cfg *config.Configs, logger *zap.SugaredLogger) Handler {
	return &handler{
		service: services,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h *handler) SetAPI(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	api := e.Group("/api/v1")
	{
		account := api.Group("/account")
		{
			account.POST("", h.CreateAccount)
			account.GET("", h.GetAllAccounts)
			account.GET("/:id", h.GetAccountByID)
			account.PUT("/:id", h.UpdateAccount)
			account.DELETE("/:id", h.DeleteAccount)
		}
		transaction := api.Group("/transaction")
		{
			transaction.POST("", h.CreateTransaction)
			transaction.GET("/account/:id", h.GetAllTransactionsByAccountID)
			transaction.GET("/:id", h.GetTransactionByID)
			transaction.DELETE("/:id", h.DeleteTransaction)
		}
	}
}

func HandleEcho(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if appErr := apperror.AsErrorInfo(err); appErr != nil {
		// очищаем DeveloperMessage, чтобы поле не присутствовало в теле http ответа
		appErr.DeveloperMessage = ""

		return c.JSON(appErr.Status, appErr)
	}

	return c.JSON(http.StatusInternalServerError, err)
}

func (h *handler) context(c echo.Context) (context.Context, context.CancelFunc) {
	ctx := c.Request().Context()

	return context.WithTimeout(ctx, h.cfg.App.Timeout)
}
