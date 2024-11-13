package repositories

import (
	"expense-tracker/models/userModels"

	"gorm.io/gorm"
)

type IAuthRepository interface {
    UserExists(username, email string) (bool, error)
    CreateUser(user userModels.User) error
}

type AuthRepository struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
    return &AuthRepository{
        db: db,
    }
}

func (r *AuthRepository) UserExists(username, email string) (bool, error) {
	var count int64

    err := r.db.Model(&userModels.User{}).
        Where("username = ? AND email = ?", username, email).
        Count(&count).
        Error

    if err != nil {
        return false, err 
    }

    return count > 0, nil
}

func (r *AuthRepository) CreateUser(user userModels.User) error {
    return r.db.Create(&user).Error
}
