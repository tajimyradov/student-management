package client

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"student-management/internal/config"
	"student-management/internal/repositories/client"
	"student-management/pkg/auth"
	"time"
)

type AuthService struct {
	repo         *client.AuthRepository
	logger       *zap.Logger
	tokenManager *auth.AuthenticationManager
	config       *config.AppConfig
}

func NewAuthService(repo *client.AuthRepository, logger *zap.Logger, tokenManager *auth.AuthenticationManager, config *config.AppConfig) *AuthService {
	return &AuthService{
		repo:         repo,
		logger:       logger,
		tokenManager: tokenManager,
		config:       config,
	}
}

func (a *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(a.config.Secrets.AccessSecret)))
}

func (a *AuthService) SignIn(username string, password string) (string, int, string, string, string, error) {
	password = a.generatePasswordHash(password)
	var userID, roleID int
	var token, firstName, lastName, image string
	//var isStudent = false
	var err error
	userID, firstName, lastName, image, err = a.repo.LoginAsStudent(username, password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			//isStudent = true
			userID, roleID, firstName, lastName, image, err = a.repo.LoginAsTeacher(username, password)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					return "", 0, "", "", "", errors.New(`Not found`)
				}
				a.logger.Info("Login as Teacher", zap.Error(err))
				return "", 0, "", "", "", errors.New("system error")
			}
		} else {
			a.logger.Info("Login As Student failed", zap.Error(err))
			return "", 0, "", "", "", errors.New("system error")
		}
	}

	roleID++

	token, err = a.tokenManager.NewJWT(userID, roleID, time.Hour*2400)
	if err != nil {
		return "", 0, "", "", "", errors.Wrap(err, "NewJWT failed")
	}

	return token, roleID, firstName, lastName, image, nil
}
