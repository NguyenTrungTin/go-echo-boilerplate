package lib

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Logger is the middleware that use to log message to the stdout
func Logger() echo.MiddlewareFunc {
	c := middleware.LoggerConfig{
		Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	}
	return middleware.LoggerWithConfig(c)
}
