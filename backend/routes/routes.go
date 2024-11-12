package routes

import (
	"expense-tracker/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler handlers.IAuthHandler) {

    userGroup := app.Group("/user")
    userGroup.Post("/register", authHandler.Register)
    

}
