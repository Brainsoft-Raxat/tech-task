package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/Brainsoft-Raxat/tech-task/internal/app/config"
	"github.com/Brainsoft-Raxat/tech-task/internal/app/connection"
	handler "github.com/Brainsoft-Raxat/tech-task/internal/handler/http"
	"github.com/Brainsoft-Raxat/tech-task/internal/repository"
	"github.com/Brainsoft-Raxat/tech-task/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Run() error {
	logger, _ := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))

	defer func() {
		_ = logger.Sync()
	}()
	
	sugar := logger.Sugar()

	cfg, err := config.New()
	if err != nil {
		sugar.Errorf("error initializing config: %v", err)
		return err
	}

	conn, err := connection.New(cfg)
	if err != nil {
		sugar.Errorf("error initializing connections: %v", err)
		return err
	}

	defer conn.Close()

	repos := repository.New(conn, cfg, sugar)
	services := service.New(repos, cfg, sugar)
	handlers := handler.New(services, cfg, sugar)

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	handlers.SetAPI(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(cfg.App.Host + ":" + cfg.App.Port); err != nil && err != http.ErrServerClosed {
			sugar.Errorf("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), cfg.App.Timeout)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		sugar.Errorf("server forced to shutdown: %v", err)
	}

	return nil
}
