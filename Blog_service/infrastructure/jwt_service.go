package infrastructure

import (
	"auth-service/config"
	"auth-service/domain"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// TokenGenerator implementation
type TokenGeneratorImpl struct{}

// NewTokenGenerator creates a new TokenGenerator instance
func NewTokenGenerator() domain.TokenGenerator {
	return &TokenGeneratorImpl{}
}

// GenerateToken generates an access token for the user
func (tg *TokenGeneratorImpl) GenerateToken(user domain.User) (string, error) {
	accessTokenSecret := []byte("TekluMoges")
	accessTokenExpiryHour := 20000

	claims := domain.JwtCustomClaims{
		UserID:    fmt.Sprintf("%d", user.ID),
		Role:     user.Role,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(accessTokenExpiryHour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(accessTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GenerateRefreshToken generates a refresh token for the user
func (tg *TokenGeneratorImpl) GenerateRefreshToken(user domain.User) (string, error) {
	refreshTokenSecret := []byte("TekluMoges")
	refreshTokenExpiryHour := 20000	

	claims := domain.JwtCustomClaims{
		UserID:   fmt.Sprintf("%d", user.ID),
		Role:     user.Role,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(refreshTokenExpiryHour))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(refreshTokenSecret)
	if err != nil {
		return "", err
	}
	return t, nil
}

// RefreshToken parses and verifies a refresh token and returns the user ID
func (tg *TokenGeneratorImpl) RefreshToken(tokenString string) (string, error) {
	refreshTokenSecret := []byte(config.EnvConfigs.JwtRefreshSecret) // Convert secret to byte slice

	t, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshTokenSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := t.Claims.(jwt.MapClaims) // Ensure type assertion to jwt.MapClaims
	if !ok || !t.Valid {
		return "", err
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", err
	}

	return userID, nil
}
