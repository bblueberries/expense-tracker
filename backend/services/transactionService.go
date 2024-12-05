package services

import (
	"errors"
	"time"

	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"github.com/bblueberries/expense-tracker/backend/repositories"
)

type ITransactionService interface {
    AddTransaction(transactionReq transactionModels.TransactionRequest) error
    DeleteTransaction(transactionID string) error
}

type TransactionService struct {
    transactionRepo repositories.ITransactionRepository
}

func NewTransactionService(transactionRepo repositories.ITransactionRepository) ITransactionService {
    return &TransactionService{
        transactionRepo: transactionRepo,
    }
}

func (s *TransactionService) AddTransaction(transactionReq transactionModels.TransactionRequest) error {
	// Validate transaction type
    if transactionReq.Type != "income" && transactionReq.Type != "expense" {
        return errors.New("invalid transaction type, must be 'income' or 'expense'")
    }

    // Validate amount
    if transactionReq.Amount <= 0 {
        return errors.New("amount must be greater than 0")
    }

    // Validate date
    if transactionReq.Date.IsZero() {
        transactionReq.Date = time.Now() 
    }

    transaction := transactionModels.Transaction{
        UserID: transactionReq.UserID,
        Type:   transactionReq.Type,
        Amount: transactionReq.Amount,
        Date:   transactionReq.Date,
		Description: transactionReq.Description,
    }

    return s.transactionRepo.CreateTransaction(transaction)
}
func (s *TransactionService) DeleteTransaction(transactionID string) error {
    // Check if transaction exists
    exists, err := s.transactionRepo.TransactionExists(transactionID)
    if err != nil {
        return errors.New("failed to check transaction existence")
    }
    if !exists {
        return errors.New("transaction not found")
    }

    // Perform deletion
    return s.transactionRepo.DeleteTransaction(transactionID)
}


