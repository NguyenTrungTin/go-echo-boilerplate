package status

import (
	"net/http"

	"github.com/labstack/echo"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from PAVE!")
}
