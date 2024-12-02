package transactionModels

import (
	"time"
)

type Transaction struct {
	ID          uint              `gorm:"primaryKey"`
	UserID      string            `gorm:"not null"`
	Type        string            `gorm:"type:varchar(20);not null"` // 'income' or 'expense'
	Amount      float64           `gorm:"not null"`
	Description string            `gorm:"type:text"`
	Date        time.Time         `gorm:"not null"`
}
