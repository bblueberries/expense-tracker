package services

import (
	"errors"
	"time"

	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"github.com/bblueberries/expense-tracker/backend/repositories"
)

type ITransactionService interface {
    AddTransaction(transactionReq transactionModels.TransactionRequest) error
    DeleteTransaction(transactionID string,userID string) error
    UpdateTransaction(transactionID string, userID string, transactionReq transactionModels.TransactionRequest) error
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
func (s *TransactionService) DeleteTransaction(transactionID string, userID string) error {
    // Check if transaction exists
    exists, err := s.transactionRepo.TransactionExists(transactionID,userID)
    if err != nil {
        return errors.New("failed to check transaction existence")
    }
    if !exists {
        return errors.New("transaction not found or not your transactions")
    }

    // Perform deletion
    err = s.transactionRepo.DeleteTransaction(transactionID)
    if err!= nil {
        return errors.New("failed to delete Transaction")
    }
    return nil
}

func (s *TransactionService) UpdateTransaction(transactionID string, userID string, transactionReq transactionModels.TransactionRequest) error {
    // Check if the transaction exists and belongs to the user
    exists, err := s.transactionRepo.TransactionExists(transactionID, userID)
    if err != nil {
        return errors.New("failed to check transaction existence")
    }
    if !exists {
        return errors.New("transaction not found or not your transaction")
    }

    updatedTransaction := transactionModels.Transaction{
        Type:        transactionReq.Type,
        Amount:      transactionReq.Amount,
        Description: transactionReq.Description,
        Date:        transactionReq.Date,
    }

    // Perform the update
    err = s.transactionRepo.UpdateTransaction(transactionID, updatedTransaction)
    if err != nil {
        return errors.New("failed to update transaction")
    }
    return nil
}


