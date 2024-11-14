package services

import (
	"errors"

	"github.com/bblueberries/expense-tracker/backend/models/userModels"
	"github.com/bblueberries/expense-tracker/backend/repositories"
	"github.com/bblueberries/expense-tracker/backend/utils"
)

type IAuthService interface {
    RegisterUser(registerReq userModels.RegisterRequest) error
    LoginUser(input userModels.LoginRequest) (string, error)
}

type AuthService struct {
    userRepo repositories.IAuthRepository
}

func NewAuthService(userRepo repositories.IAuthRepository) IAuthService {
    return &AuthService{
        userRepo: userRepo,
    }
}

func (s *AuthService) RegisterUser(registerReq userModels.RegisterRequest) error {
	//check if user exist?
    if exists, _ := s.userRepo.UserExists(registerReq.Username); exists {
        return errors.New("user already exists")
    }

    // Hash the password
    hashedPassword, err := utils.HashPassword(registerReq.Password)
    if err != nil {
        return errors.New("failed to hash password")
    }

    user := userModels.User{
        UserID:       registerReq.Username, 
        Username:     registerReq.Username,
        PasswordHash: hashedPassword,
        Email:        registerReq.Email,
    }

    return s.userRepo.CreateUser(user)
}

func (s *AuthService) LoginUser(input userModels.LoginRequest) (string, error) {
    // get user by username
    user, err := s.userRepo.GetUserByUsername(input.Username)
    if err != nil {
        return "", errors.New("invalid username or password")
    }

    // check hash
    if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
        return "", errors.New("invalid username or password")
    }

    // generate JWT
    token, err := utils.GenerateJWTToken(user)
    if err != nil {
        return "", errors.New("could not generate authentication token")
    }

    return token, nil
}

