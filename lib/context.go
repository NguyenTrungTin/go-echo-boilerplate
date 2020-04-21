package lib

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
}

func CustomContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	}
}
