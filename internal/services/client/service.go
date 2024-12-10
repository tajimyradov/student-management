package client

import (
	"go.uber.org/zap"
	"student-management/internal/config"
	"student-management/internal/repositories/client"
	"student-management/pkg/auth"
)

type Service struct {
	AuthService      *AuthService
	TimetableService *TimetableService
	StudentService   *StudentService
}

type ServiceDeps struct {
	Repos *client.Repository

	Logger       *zap.Logger
	Config       *config.AppConfig
	TokenManager *auth.AuthenticationManager
}

func NewService(deps *ServiceDeps) *Service {
	return &Service{
		AuthService:      NewAuthService(deps.Repos.AuthRepository, deps.Logger, deps.TokenManager, deps.Config),
		TimetableService: NewTimetableService(deps.Repos.TimetableRepository, deps.Logger),
		StudentService:   NewStudentService(deps.Repos.StudentRepository, deps.Logger),
	}
}
