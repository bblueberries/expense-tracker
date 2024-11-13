package handlers

import (
	"log"

	"github.com/bblueberries/expense-tracker/backend/models/userModels"
	"github.com/bblueberries/expense-tracker/backend/response"
	"github.com/bblueberries/expense-tracker/backend/services"

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
        return response.Error(c,fiber.StatusBadRequest,"Invalid request body")
    }
    

    err := h.AuthService.RegisterUser(registerReq)
    if err != nil {
        log.Println("Error registering user:", err)
        return response.Error(c,fiber.StatusBadRequest,"Failed to register user")
    }

    return response.Success(c,fiber.StatusCreated,"user registered successfully!",nil)
   
}
