package database

import (
	"log"

	"github.com/bblueberries/expense-tracker/backend/models/transactionModels"
	"github.com/bblueberries/expense-tracker/backend/models/userModels"
	"gorm.io/gorm"
)


func RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&userModels.User{},
		&transactionModels.Transaction{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("All migrations ran successfully")
}
