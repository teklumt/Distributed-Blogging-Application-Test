package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Token represents the structure of the token entity in MySQL
type Token struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"` // Use GORM for MySQL
	UserID       string      `json:"userId"`                              // References the User ID
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// JwtCustomClaims represents the custom claims for JWT
type JwtCustomClaims struct {
	UserID     string   `json:"user_id"`
	Role       string `json:"role"`
	Username   string `json:"username"`
	jwt.RegisteredClaims
}
