// repository/user_repository.go
package repository

import (
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/config"
	"github.com/teklumt/Distributed-Blogging-Application-Test-auth-service/domain"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type userRepository struct{}

// Initialize a new user repository
func NewAuthRepository() domain.UserRepository {
	return &userRepository{}
}

// Register a new user and publish a registration event to RabbitMQ
func (ur *userRepository) Register(user domain.User) error {
	// Check if user already exists
	_, err := ur.GetUserByEmail(user.Email)
	if err == nil {
		return errors.New("User already exists")
	}
	user.TokenCreatedAt = nil

	// Save the new user to the database
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	// Publish user registration event to RabbitMQ
	// newUser := domain.User{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// }

	// err = messaging.PublishUserRegistrationEvent(messaging.UserRegistrationEvent{
	// 	UserID:   newUser.ID,
	// 	Username: newUser.Username,
	// })
	// if err != nil {
	// 	log.Printf("Error publishing registration event: %v", err)
	// 	return err
	// }

	// log.Println("User registered and event published successfully")
	return nil
}

// Fetch a user by email
func (ur *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

// Authenticate a user by email and password
// func (ur *userRepository) Login(email, password string) (domain.User, error) {
// 	var user domain.User
// 	err := config.DB.Where("email = ?", email).First(&user).Error
// 	return user, err
// }

func (ur *userRepository) Login(email, password string) (domain.User, error) {
	var user domain.User
	// Find user by email
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return domain.User{}, errors.New("user not found")
	}

	// Compare provided password with stored password (assuming password is hashed)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, errors.New("incorrect password")
	}

	// Publish login event to RabbitMQ
	// err = messaging.PublishUserLoginEvent(messaging.UserLoginEvent{
	// 	UserID:   user.ID,
	// 	Username: user.Username,
	// })
	// if err != nil {
	// 	log.Printf("Error publishing login event: %v", err)
	// 	return domain.User{}, errors.New("failed to process login")
	// }

	return user, nil
}
