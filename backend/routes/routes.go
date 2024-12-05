package routes

import (
	"github.com/bblueberries/expense-tracker/backend/handlers"
	"github.com/bblueberries/expense-tracker/backend/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler handlers.IAuthHandler,transactionHandler handlers.ITransactionHandler) {

    userGroup := app.Group("/user")
    userGroup.Post("/register", authHandler.Register)
    userGroup.Post("/login",authHandler.Login)


    transactionGroup := app.Group("/transaction").Use(middlewares.AuthorizationUserToken())
    transactionGroup.Post("",transactionHandler.AddTransaction)
    transactionGroup.Delete("/:id",transactionHandler.DeleteTransaction)
    transactionGroup.Put("/:id",transactionHandler.UpdateTransaction)

    
}
