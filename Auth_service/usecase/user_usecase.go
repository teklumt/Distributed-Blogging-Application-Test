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

func (uc *UserUseCase) Register(user domain.User) domain.ErrorResponse {
	check, _ := uc.UserRepository.GetUserByEmail(user.Email)
	if check.Email != "" {
		return domain.ErrorResponse{Message: "Email already exists", StatusCode: 400}
	}

	user.Role = "user"

	hashedPassword, err := uc.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.ErrorResponse{Message: "Failed to hash password", StatusCode: 500}
	}
	user.Password = hashedPassword

	err = uc.UserRepository.Register(user)
	if err != nil {
		return domain.ErrorResponse{Message: "Failed to create user", StatusCode: 500}
	}

	return domain.ErrorResponse{}
}

func (uc *UserUseCase) Login(email string, password string) (domain.LogInResponse, domain.ErrorResponse) {
	user, err := uc.UserRepository.GetUserByEmail(email)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrorResponse{Message: "User not found", StatusCode: 404}
	}

	if !uc.PasswordSvc.CheckPasswordHash(password, user.Password) {
		return domain.LogInResponse{}, domain.ErrorResponse{Message: "Invalid credentials", StatusCode: 400}
	}

	token, err := uc.TokenGen.GenerateToken(user)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrorResponse{Message: "Failed to generate token", StatusCode: 500}
	}

	refreshToken, err := uc.TokenGen.GenerateRefreshToken(user)
	if err != nil {
		return domain.LogInResponse{}, domain.ErrorResponse{Message: "Failed to generate refresh token", StatusCode: 500}
	}


	return domain.LogInResponse{
		AccessToken: token,
		RefreshToken: refreshToken,
		User: domain.ReturnUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
		},
	}, domain.ErrorResponse{}
}
