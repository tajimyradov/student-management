package services

import (
	"go.uber.org/zap"
	"student-management/internal/config"
	repository "student-management/internal/repositories"
	"student-management/internal/services/admin"
	"student-management/internal/services/client"
	"student-management/pkg/auth"
)

type ServiceDeps struct {
	Repos *repository.Repositories

	Logger       *zap.Logger
	Config       *config.AppConfig
	TokenManager *auth.AuthenticationManager
}

type Service struct {
	AdminService  *admin.Service
	ClientService *client.Service
}

func NewServices(deps *ServiceDeps) *Service {
	return &Service{
		AdminService: admin.NewService(&admin.ServiceDeps{
			Repos:  deps.Repos.AdminRepos,
			Logger: deps.Logger,
			Config: deps.Config,
		}),

		ClientService: client.NewService(&client.ServiceDeps{
			Repos:        deps.Repos.ClientRepos,
			Logger:       deps.Logger,
			Config:       deps.Config,
			TokenManager: deps.TokenManager,
		}),
	}
}
