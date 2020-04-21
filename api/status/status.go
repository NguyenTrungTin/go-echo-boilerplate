package status

import (
	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/healthz", Healthz)
}
