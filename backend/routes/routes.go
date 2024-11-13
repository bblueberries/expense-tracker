package routes

import (
	"github.com/bblueberries/expense-tracker/backend/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler handlers.IAuthHandler) {

    userGroup := app.Group("/user")
    userGroup.Post("/register", authHandler.Register)
    

}
