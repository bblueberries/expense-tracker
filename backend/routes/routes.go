package routes

import (
	"github.com/bblueberries/expense-tracker/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler handlers.IAuthHandler,transactionHandler handlers.ITransactionHandler) {

    userGroup := app.Group("/user")
    userGroup.Post("/register", authHandler.Register)
    userGroup.Post("/login",authHandler.Login)


    transactionGroup := app.Group("/transaction")
    transactionGroup.Post("",transactionHandler.AddTransaction)

    
}
