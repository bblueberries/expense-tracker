package middlewares

import (
	"os"

	"github.com/bblueberries/expense-tracker/backend/response"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func AuthorizationUserToken() fiber.Handler {
    return jwtware.New(jwtware.Config{
        SigningKey:   []byte(os.Getenv("JWT_SECRET")),
        ErrorHandler: AuthError, 
        SuccessHandler: AuthSuccess, 
    })
}


func AuthError(c *fiber.Ctx, err error) error {
    return response.Error(c,fiber.StatusBadRequest,"Unauthorized")
   
}
func AuthSuccess(c *fiber.Ctx) error { 
    c.Locals("authSuccess", true) 
    return c.Next()              
}
