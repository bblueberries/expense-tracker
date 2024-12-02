package handlers

import (
	"log"

	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"github.com/bblueberries/expense-tracker/backend/response"
	"github.com/bblueberries/expense-tracker/backend/services"

	"github.com/gofiber/fiber/v2"
)

type ITransactionHandler interface {
    AddTransaction(c *fiber.Ctx) error
}

type TransactionHandler struct {
    TransactionService services.ITransactionService
}

func NewTransactionHandler(transactionService services.ITransactionService) ITransactionHandler {
    return &TransactionHandler{
        TransactionService: transactionService,
    }
}

func (h *TransactionHandler) AddTransaction(c *fiber.Ctx) error {
    var transactionReq transactionModels.TransactionRequest

    if err := c.BodyParser(&transactionReq); err != nil {
        return response.Error(c, fiber.StatusBadRequest, "Invalid request body")
    }

    // Call service to add the transaction
    if err := h.TransactionService.AddTransaction(transactionReq); err != nil {
        log.Println("Error adding transaction:", err)
        return response.Error(c, fiber.StatusBadRequest, "Failed to creating transaction")
    }

    return response.Success(c, fiber.StatusCreated, "Transaction added successfully!", nil)
}