package handlers

import (
	"log"

	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"github.com/bblueberries/expense-tracker/backend/response"
	"github.com/bblueberries/expense-tracker/backend/services"
	"github.com/bblueberries/expense-tracker/backend/utils"

	"github.com/gofiber/fiber/v2"
)

type ITransactionHandler interface {
    AddTransaction(c *fiber.Ctx) error
    DeleteTransaction(c *fiber.Ctx) error
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
	// Get user_id from the token
	userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token")
	}
	transactionReq.UserID = userID
    // Call service to add the transaction
    if err := h.TransactionService.AddTransaction(transactionReq); err != nil {
        log.Println("Error adding transaction:", err)
        return response.Error(c, fiber.StatusBadRequest, "Failed to creating transaction")
    }

    return response.Success(c, fiber.StatusCreated, "Transaction added successfully!", nil)
}

func (h *TransactionHandler) DeleteTransaction(c *fiber.Ctx) error {
    id := c.Params("id") // Get transaction ID from URL
    if id == "" {
        return response.Error(c, fiber.StatusBadRequest, "Transaction ID is required")
    }
    userID, err := utils.GetUserIDFromToken(c)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token")
	}

    err = h.TransactionService.DeleteTransaction(id,userID)
    if err != nil {
        return response.Error(c, fiber.StatusInternalServerError, err.Error())
    }

    return response.Success(c, fiber.StatusOK, "Transaction deleted successfully", nil)
}