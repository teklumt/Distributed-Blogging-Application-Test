// repository/user_repository.go
package repository

import (
	"auth-service/config"
	"auth-service/domain"
)

type userRepository struct{}

// Initialize a new user repository
func NewAuthRepository() domain.UserRepository {
	return &userRepository{}
}



// Fetch a user by email
func (ur *userRepository) GetUserByEmail(email string) (domain.User, error) {
	var user domain.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}




func (ur *userRepository) GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	return user, err
}


// Update a user
func (ur *userRepository) UpdateUser(id uint, user domain.User) (domain.User, error) {
	// Update the user in the database
	err := config.DB.Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return domain.User{}, err
	}

	// Fetch the updated user
	updatedUser, err := ur.GetUserByID(id)
	if err != nil {
		return domain.User{}, err
	}

	// // Publish account update event to RabbitMQ

	// err = messaging.PublishUserAccountUpdateEvent(messaging.UserAccountUpdateEvent{
	// 	UserID:   updatedUser.ID,
	// 	Username: updatedUser.Username,
	// })

	// if err != nil {
	// 	return domain.User{}, err
	// }

	

	return updatedUser, nil
	
}