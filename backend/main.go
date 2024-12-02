package main

import (
	"log"

	"github.com/bblueberries/expense-tracker/backend/database"
	"github.com/bblueberries/expense-tracker/backend/handlers"
	"github.com/bblueberries/expense-tracker/backend/repositories"
	"github.com/bblueberries/expense-tracker/backend/routes"
	"github.com/bblueberries/expense-tracker/backend/services"

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

    transactionRepo := repositories.NewTransactionRepository(db)
    transactionService := services.NewTransactionService(transactionRepo)
    transactionHandler := handlers.NewTransactionHandler(transactionService)


    routes.SetupRoutes(app, authHandler,transactionHandler)


    log.Fatal(app.Listen(":3000"))
}



