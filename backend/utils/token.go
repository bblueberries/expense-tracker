// utils/token.go
package utils

import (
	"os"
	"time"

	"github.com/bblueberries/expense-tracker/backend/models/userModels"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWTToken(user *userModels.User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.UserID,
        "exp":     time.Now().Add(1 * time.Hour).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    secretKey := os.Getenv("JWT_SECRET")
    return token.SignedString([]byte(secretKey))
}
