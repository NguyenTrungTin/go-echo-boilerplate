package lib

import "github.com/labstack/echo"

type Resp struct {
	Message string `json:"message"`
}

func JSON(code int, message string) *echo.HTTPError {
	return echo.NewHTTPError(code, message)
}
