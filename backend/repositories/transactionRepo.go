package repositories

import (
	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
    CreateTransaction(transaction transactionModels.Transaction) error
    TransactionExists(transactionID string) (bool, error)
    DeleteTransaction(transactionID string) error
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

func (r *TransactionRepository) TransactionExists(transactionID string) (bool, error) {
    var count int64
    err := r.db.Model(&transactionModels.Transaction{}).Where("id = ?", transactionID).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *TransactionRepository) DeleteTransaction(transactionID string) error {
    return r.db.Delete(&transactionModels.Transaction{}, "id = ?", transactionID).Error
}
