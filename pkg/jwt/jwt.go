package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func CreateToken(userId uint64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(os.Getenv("JWT_KEY")))
}
