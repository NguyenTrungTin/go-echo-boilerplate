package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetUserID(c echo.Context) (uint, bool) {
	u, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return 0, false
	}
	claims, ok := u.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}
	id, ok := claims["id"].(float64)
	if !ok {
		return 0, false
	}
	return uint(id), true
}
