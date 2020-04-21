package auth

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
)

// PublicPath is a slice contain all  paths that we don't use JWT middleware
var PublicPath = []string{"/", "/healthz", "/api/user/register", "/api/user/login", "/api/dealership/template"}

var StartBy = []string{"/api/inspection-result"}

// JWT provides a JSON Web Token (JWT) authentication middleware.
//
// For valid token, it sets the user in context and calls next handler.
// For invalid token, it sends 401 - Unauthorized response.
// For missing or invalid Authorization header, it sends 400 - Bad Request.
func JWT() echo.MiddlewareFunc {
	c := middleware.JWTConfig{
		SigningKey:  []byte(config.JWT_KEY),
		TokenLookup: "header:Authorization",
		Skipper: func(c echo.Context) bool {
			if _, ok := lib.Find(PublicPath, c.Request().URL.Path); ok == true {
				return true
			}

			for _, item := range StartBy {
				if strings.HasPrefix(c.Request().URL.Path, item) {
					return true
				}
			}

			return false
		},
	}

	return middleware.JWTWithConfig(c)
}
