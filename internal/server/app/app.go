package server

import (
	"context"
	"fmt"
	"github.com/vleukhin/GophKeeper/internal/server/api/v1"
	"github.com/vleukhin/GophKeeper/internal/server/storage/postgres"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	config "github.com/vleukhin/GophKeeper/internal/config/server"
	"github.com/vleukhin/GophKeeper/internal/pkg/httpserver"
	"github.com/vleukhin/GophKeeper/internal/pkg/logger"
	"github.com/vleukhin/GophKeeper/internal/server/core"
)

func Run(ctx context.Context, cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	storage, err := postgres.NewPostgresStorage(cfg.Postgres.DSN, time.Second*5)
	if err != nil {
		l.Fatal(err.Error())
	}
	err = storage.Migrate(ctx)
	if err != nil {
		l.Fatal(err.Error())
	}

	app := core.New(
		storage,
		cfg,
		l,
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, app, l)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
