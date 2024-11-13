package response

import (
	"github.com/bblueberries/expense-tracker/backend/models"

	"github.com/gofiber/fiber/v2"
)
func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
    return c.Status(statusCode).JSON(models.Response{
        StatusCode: statusCode,
        Message:    message,
        Data:       data,
    })
}

func Error(c *fiber.Ctx, statusCode int, message string) error {
    return c.Status(statusCode).JSON(models.Response{
        StatusCode: statusCode,
        Message:    message,
    })
}