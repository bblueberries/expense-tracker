package main

import (
	"expense-tracker/database"
	"expense-tracker/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	// Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    // Initialize Fiber app
    app := fiber.New()

    db, err := database.ConnectDB()
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    //Insert a new user
    err = insertUser(db, "testuser")
    if err != nil {
        log.Fatalf("Failed to insert user: %v", err)
    }

    // Retrieve and print all users
    users, err := getAllUsers(db)
    if err != nil {
        log.Fatalf("Failed to retrieve users: %v", err)
    }

    fmt.Println("Users in 'test' table:")
    for _, user := range users {
        fmt.Printf("ID: %d, User: %s\n", user.ID, user.User)
    }
    // Register routes
    // routes.RegisterRoutes(app, db)

    // Start the server
    log.Fatal(app.Listen(":3000"))
}


// Insert a new user into the 'test' table
func insertUser(db *gorm.DB, username string) error {
    user := models.Test{User: username}
    return db.Create(&user).Error
}

// Retrieve all users from the 'test' table
func getAllUsers(db *gorm.DB) ([]models.Test, error) {
    var users []models.Test
    err := db.Table("test").Find(&users).Error
    return users, err
}