package database

import (
	"expense-tracker/models/userModels"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.AutoMigrate(&userModels.User{})
    if err != nil {
        return nil, fmt.Errorf("failed to auto-migrate models: %w", err)
    }
    return db, nil
}
