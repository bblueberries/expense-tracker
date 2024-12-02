package repositories

import (
	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
    CreateTransaction(transaction transactionModels.Transaction) error
}


type TransactionRepository struct {
    db *gorm.DB
}


func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
    return &TransactionRepository{
        db: db,
    }
}


func (r *TransactionRepository) CreateTransaction(transaction transactionModels.Transaction) error {
    return r.db.Create(&transaction).Error
}
