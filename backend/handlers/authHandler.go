package handlers

import (
	"expense-tracker/models/userModels"
	"expense-tracker/services"
	"log"

	"github.com/gofiber/fiber/v2"
)
type IAuthHandler interface {
    Register(c *fiber.Ctx) error
}

type AuthHandler struct {
    AuthService services.IAuthService
}


func NewAuthHandler(authService services.IAuthService) IAuthHandler {
    return &AuthHandler{
        AuthService: authService,
    }
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var registerReq userModels.RegisterRequest

    // handle data 
    if err := c.BodyParser(&registerReq); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request body",
        })
    }

    err := h.AuthService.RegisterUser(registerReq)
    if err != nil {
        log.Println("Error registering user:", err)
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to register user",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "User registered successfully",
    })
}
