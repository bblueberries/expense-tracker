package repositories

import (
	"github.com/bblueberries/expense-tracker/backend/models/userModels"
	"gorm.io/gorm"
)

type IAuthRepository interface {
    UserExists(username string) (bool, error)
    CreateUser(user userModels.User) error
    GetUserByUsername(username string) (*userModels.User, error)
}

type AuthRepository struct {
    db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) IAuthRepository {
    return &AuthRepository{
        db: db,
    }
}

func (r *AuthRepository) UserExists(username string) (bool, error) {
	var count int64

    //in case username and email, individual is not unique, (rn username unique) 
    err := r.db.Model(&userModels.User{}).
        Where("username = ?", username).
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

func (r *AuthRepository) GetUserByUsername(username string) (*userModels.User, error) {
    var user userModels.User
    err := r.db.Model(&userModels.User{}).
    Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
