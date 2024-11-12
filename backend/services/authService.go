package services

import (
	"errors"
	"expense-tracker/models"
	"expense-tracker/repositories"

	"expense-tracker/utils"
)

type IAuthService interface {
    RegisterUser(registerReq models.RegisterRequest) error
}

type AuthService struct {
    userRepo repositories.IAuthRepository
}

func NewAuthService(userRepo repositories.IAuthRepository) IAuthService {
    return &AuthService{
        userRepo: userRepo,
    }
}

func (s *AuthService) RegisterUser(registerReq models.RegisterRequest) error {
	//check if user exist?
    if exists, _ := s.userRepo.UserExists(registerReq.Username, registerReq.Email); exists {
        return errors.New("user already exists")
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(registerReq.Password)
    if err != nil {
        return errors.New("failed to hash password")
    }

    user := models.User{
        UserID:       registerReq.Username, 
        Username:     registerReq.Username,
        PasswordHash: hashedPassword,
        Email:        registerReq.Email,
    }

    return s.userRepo.CreateUser(user)
}

