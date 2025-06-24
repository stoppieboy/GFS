package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(username, password string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type authService struct {
	secret string
}

func NewAuthService(secret string) AuthService {
	return &authService{secret: secret}
}

func (a *authService) Login(username, password string) (string, error) {
	if username != "admin" || password != "password" {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(a.secret))
}

func (a *authService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.secret), nil
	})
}