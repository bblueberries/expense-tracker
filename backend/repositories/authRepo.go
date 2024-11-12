package repositories

import (
	"expense-tracker/models"

	"gorm.io/gorm"
)

type IAuthRepository interface {
    UserExists(username, email string) (bool, error)
    CreateUser(user models.User) error
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

    err := r.db.Model(&models.User{}).
        Where("username = ? OR email = ?", username, email).
        Count(&count).
        Error

    if err != nil {
        return false, err 
    }

    return count > 0, nil
}

func (r *AuthRepository) CreateUser(user models.User) error {
    return r.db.Create(&user).Error
}
