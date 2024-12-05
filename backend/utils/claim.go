package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromToken(c *fiber.Ctx) (string, error) {

	userToken := c.Locals("user").(*jwt.Token)
	if userToken == nil {
		return "", fiber.NewError(fiber.StatusUnauthorized, "No token found")
	}

	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok || !userToken.Valid {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid or malformed token")
	}
	
	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", fiber.NewError(fiber.StatusUnauthorized, "Invalid token payload: user_id missing")
	}

	return userID, nil
}