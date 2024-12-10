package server

import (
	"context"
	"net/http"
	"student-management/internal/app/server/handlers"
	"student-management/internal/config"
	repository "student-management/internal/repositories"
	"student-management/internal/services"
	"student-management/pkg/auth"

	errch "github.com/proxeter/errors-channel"
	"go.uber.org/zap"
)

type Server struct {
	ctx context.Context

	logger *zap.Logger
	config *config.AppConfig

	repos   *repository.Repositories
	handler *handlers.Handler
}

func NewServer(
	ctx context.Context,
	logger *zap.Logger,
	config *config.AppConfig,
	repos *repository.Repositories,
	tokenManager *auth.AuthenticationManager,

) <-chan error {

	return errch.Register(func() error {
		srcDeps := &services.ServiceDeps{
			Repos:        repos,
			Logger:       logger,
			Config:       config,
			TokenManager: tokenManager,
		}

		src := services.NewServices(srcDeps)

		handler := handlers.NewHandler(logger, config, src)

		return (&Server{
			ctx: ctx,

			config: config,
			logger: logger,

			handler: handler,
		}).Start(ctx)
	})
}

func (s *Server) Start(ctx context.Context) error {
	h := s.handler.Init()

	server := http.Server{
		Handler: h,
		Addr:    ":" + s.config.HTTP.Port,
	}

	s.logger.Info(
		"Server running",
		zap.String("host", s.config.HTTP.Host),
		zap.String("port", s.config.HTTP.Port),
	)

	select {
	case err := <-errch.Register(server.ListenAndServe):
		s.logger.Info("Shutdown api server", zap.String("by", "error"), zap.Error(err))
		return server.Shutdown(ctx)
	case <-ctx.Done():
		s.logger.Info("Shutdown api server", zap.String("by", "context.Done"))
		return server.Shutdown(ctx)
	}
}
