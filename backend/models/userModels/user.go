package userModels

import "github.com/bblueberries/expense-tracker/backend/models/transactionModels"

type User struct {
    UserID       string `gorm:"primaryKey" json:"user_id"` 
    Username     string `gorm:"unique;not null" json:"username"`
    PasswordHash string `gorm:"not null" json:"password_hash"`
    Email        string `gorm:"not null" json:"email"`

    Transactions []transactionModels.Transaction //`gorm:"foreignKey:UserID;references:UserID"`
}

