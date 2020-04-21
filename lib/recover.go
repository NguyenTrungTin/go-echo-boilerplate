package lib

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Recover is the middleware that used to recovery frm panic
func Recover() echo.MiddlewareFunc {
	c := middleware.RecoverConfig{
		StackSize: 4 << 10,
	}
	return middleware.RecoverWithConfig(c)
}
