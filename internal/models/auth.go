package models

import "github.com/dgrijalva/jwt-go"

type UserClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
}

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
