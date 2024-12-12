package transactionModels

import (
	"time"
)
type TransactionRequest struct {
    UserID		 string      `json:"user_id"`  
    Type  		 string    `json:"type"`   
    Amount       float64   `json:"amount"`
    Date   	     time.Time `json:"date"`   
	Description  string    `json:"description"`
}
