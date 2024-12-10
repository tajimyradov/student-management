package app

import (
	"context"
	"student-management/internal/app/server"
	"student-management/internal/config"
	repository "student-management/internal/repositories"
	"student-management/pkg/auth"

	"go.uber.org/zap"
)

type App struct {
	config       *config.AppConfig
	repos        *repository.Repositories
	logger       *zap.Logger
	tokenManager *auth.AuthenticationManager
}

func New(config *config.AppConfig, logger *zap.Logger, repos *repository.Repositories, tokenManager *auth.AuthenticationManager) *App {
	return &App{
		config:       config,
		logger:       logger,
		repos:        repos,
		tokenManager: tokenManager,
	}
}

func (app *App) Start(ctx context.Context) {
	httpErrChan := server.NewServer(
		ctx,
		app.logger,
		app.config,
		app.repos,
		app.tokenManager,
	)
	<-httpErrChan
}

func (app *App) Shutdown() {
	app.logger.Info("Shutdown logger")
	_ = app.logger.Sync()
}
