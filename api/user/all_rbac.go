package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
)

func GetAllRBAC(c echo.Context) error {
	resp := struct {
		Success bool     `json:"success"`
		AllRBAC []string `json:"all_rbac"`
	}{
		Success: true,
		AllRBAC: auth.RBAC.AllRBAC,
	}

	return c.JSON(http.StatusOK, &resp)
}
