package domain

import (
	"encoding/json"
	"time"
)

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
    ID                uint      `gorm:"primaryKey;autoIncrement" json:"id"`
    Username          string    `gorm:"unique;not null" json:"username"`
    Email             string    `gorm:"unique;not null" json:"email"`
    Password          string    `gorm:"column:password_hash;not null" json:"password"`
    Role              string    `gorm:"role" json:"role"`
    CreatedAt         time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"createdAt"`
    UpdatedAt         time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt"`
    Image             string    `gorm:"image,omitempty" json:"image,omitempty"`
    ActivationToken   string    `gorm:"activation_token,omitempty" json:"activation_token,omitempty"`
    TokenCreatedAt    *time.Time `gorm:"column:token_created_at;default:NULL" json:"token_created_at,omitempty"`
    IsActive          bool      `gorm:"is_active"`
    GoogleID          string    `gorm:"google_id,omitempty" json:"google_id,omitempty"`
    PasswordResetToken string   `gorm:"password_reset_token,omitempty" json:"password_reset_token,omitempty"`
}
func (User) TableName() string {
    return "Users"  
}

type ReturnUser struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"username" json:"username"`
	Email     string    `gorm:"email" json:"email"`
	Role      string    `gorm:"role" json:"role"`
	CreatedAt time.Time `gorm:"createdAt" json:"createdAt"`
	Image     string    `gorm:"image,omitempty" json:"image,omitempty"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(&struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		Role     string `json:"role"`
		Image    string `json:"image,omitempty"`
	}{
		Username: u.Username,
		Email:    u.Email,
		Name:     u.Username,
		Role:     u.Role,
		Image:    u.Image,
	})
}

type OAuthProvider string

const (
	Google OAuthProvider = "google"
)

type RefreshToken struct {
	Token     string    `gorm:"column:token" json:"token"`
	DeviceID  string    `gorm:"column:device_id" json:"device_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

type LogInResponse struct {
	AccessToken  string     `json:"accessToken"`
	RefreshToken string     `json:"refreshToken"`
	User         ReturnUser `json:"user"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type TokenGenerator interface {
	GenerateToken(user User) (string, error)
	GenerateRefreshToken(user User) (string, error)
	RefreshToken(token string) (string, error)
}

type TokenVerifier interface {
	VerifyToken(token string) (*User, error)
	VerifyRefreshToken(token string) (*User, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type UserUsecase interface {
	Register(user User) ErrorResponse
	Login(email string, password string) (LogInResponse, ErrorResponse)
}

type UserRepository interface {
	Register(user User) error
	GetUserByEmail(email string) (User, error)
	Login(email string, password string) (User, error)
}
