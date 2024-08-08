package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type UserTest struct {
	ID    string
	Name  string
	Point int
}

func testDBconnected(db *gorm.DB){
	var user UserTest
	// Retrieve the first row from the user_test table
	if err := db.Table("user_test").First(&user).Error; err != nil {
		log.Fatalf("failed to get data: %v", err)
	}
	fmt.Printf("ID: %s, Name: %s, Point: %d\n", user.ID, user.Name, user.Point)
}

func main() {
	app := fiber.New()
	
	err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	app.Use(cors.New(cors.Config{
			Next:             nil,
			AllowOriginsFunc: nil,
			AllowOrigins:     "*",
			AllowMethods: strings.Join([]string{
				fiber.MethodGet,
				fiber.MethodPost,
				fiber.MethodHead,
				fiber.MethodPut,
				fiber.MethodDelete,
				fiber.MethodPatch,
			}, ","),
			AllowHeaders:     "",
			AllowCredentials: false,
			ExposeHeaders:    "",
			MaxAge:           0,
		}))
	
	dsn := os.Getenv("DATABASE_CONNECTION_STRING")
	if dsn == "" {
		log.Fatal("DATABASE_CONNECTION_STRING environment variable is not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("Database connection successful!")

	//test get 1st row from table 'user_test'
	testDBconnected(db)
   
	app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })
	if err := app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
