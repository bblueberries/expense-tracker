package main

import (
	"log"

	"github.com/bblueberries/expense-tracker/database"
	"github.com/bblueberries/expense-tracker/handlers"
	"github.com/bblueberries/expense-tracker/repositories"
	"github.com/bblueberries/expense-tracker/routes"
	"github.com/bblueberries/expense-tracker/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    app := fiber.New()

    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }
   
    authRepo := repositories.NewAuthRepository(db)
    authService := services.NewAuthService(authRepo)
    authHandler := handlers.NewAuthHandler(authService)


    routes.SetupRoutes(app, authHandler)


    log.Fatal(app.Listen(":3000"))
}



