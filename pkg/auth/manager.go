package auth

import (
	"github.com/dgrijalva/jwt-go"
	"student-management/internal/models"
	"time"
)

type AuthenticationManager struct {
	signingKey string
}

func NewAuthenticationManager(signingKey string) *AuthenticationManager {

	return &AuthenticationManager{signingKey: signingKey}
}

func (m *AuthenticationManager) NewJWT(userId, roleId int, ttl time.Duration) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, models.UserClaims{
		UserID: userId,
		RoleID: roleId,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
		},
	})

	return token.SignedString([]byte(m.signingKey))
}
