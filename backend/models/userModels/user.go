package userModels

type User struct {
    UserID       string `gorm:"primaryKey" json:"user_id"` 
    Username     string `gorm:"unique;not null" json:"username"`
    PasswordHash string `gorm:"not null" json:"password_hash"`
    Email        string `gorm:"unique;not null" json:"email"`
}

