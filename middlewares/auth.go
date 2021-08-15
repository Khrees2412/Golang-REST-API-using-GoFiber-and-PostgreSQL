package auth

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)




func GetUserID(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRETKEY")), nil
	})

	claims, _ := token.Claims.(*jwt.StandardClaims)
	userID := claims.Issuer
	
	return userID, err
}