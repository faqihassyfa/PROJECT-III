package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SecretJWT")),
	})
}

func CreateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 504).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SecretJWT")))
}

func ExtractToken(e echo.Context) (data map[string]interface{}, err error) {
	user := e.Get("user").(*jwt.Token)
	var dataToken = map[string]interface{}{}
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		dataToken["userID"] = claims["userID"].(float64)
		dataToken["role"] = claims["role"].(string)
		return dataToken, nil
	}
	return nil, fmt.Errorf("token invalid")
}
