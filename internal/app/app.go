package app

import (
	"context"
	"student-management/internal/app/server"
	"student-management/internal/config"
	repository "student-management/internal/repositories"

	"go.uber.org/zap"
)

type App struct {
	config *config.AppConfig
	repos  *repository.Repository
	logger *zap.Logger
}

func New(config *config.AppConfig, logger *zap.Logger, repos *repository.Repository) *App {
	return &App{
		config: config,
		logger: logger,
		repos:  repos,
	}
}

func (app *App) Start(ctx context.Context) {
	httpErrChan := server.NewServer(
		ctx,
		app.logger,
		app.config,
		app.repos,
	)
	<-httpErrChan
}

func (app *App) Shutdown() {
	app.logger.Info("Shutdown logger")
	_ = app.logger.Sync()
}
