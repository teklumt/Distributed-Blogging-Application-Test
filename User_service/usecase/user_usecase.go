package usecase

import (
	"auth-service/domain"
) 





type UserUseCase struct {
	UserRepository domain.UserRepository
	TokenGen       domain.TokenGenerator
	PasswordSvc    domain.PasswordService
}

func NewAuthUsecase(userRepository domain.UserRepository, tokenGen domain.TokenGenerator, passwordSvc domain.PasswordService) domain.UserUsecase {
	return &UserUseCase{
		UserRepository: userRepository,
		TokenGen:       tokenGen,
		PasswordSvc:    passwordSvc,
	}
}




func (uc *UserUseCase) GerUserByID(id uint) (domain.ReturnUser, domain.ErrorResponse) {
	user, err := uc.UserRepository.GetUserByID(id)
	if err != nil {
		return domain.ReturnUser{}, domain.ErrorResponse{Message: "User not found", StatusCode: 404}
	}

	return domain.ReturnUser{
		ID:        user.ID,
		Email:     user.Email,
		Username: user.Username,
		Role:    user.Role,
		CreatedAt: user.CreatedAt,
		Image:   "picture",
	}, domain.ErrorResponse{}
}



func (uc *UserUseCase) UpdateUser(id uint, user domain.User) (domain.ReturnUser, domain.ErrorResponse) {
	var newUser domain.User
	_, err := uc.UserRepository.GetUserByID(id)
	if err != nil {
		return domain.ReturnUser{}, domain.ErrorResponse{Message: "User not found", StatusCode: 404}
	}

	newUser, err = uc.UserRepository.UpdateUser(id, user)
	if err != nil {
		return domain.ReturnUser{}, domain.ErrorResponse{Message: "Failed to update user", StatusCode: 500}
	}

	return domain.ReturnUser{
		ID:        newUser.ID,
		Email:     user.Email,
		Username: user.Username,
		Role:    user.Role,
		CreatedAt: user.CreatedAt,
		Image:   "picture",
	}, domain.ErrorResponse{}
}