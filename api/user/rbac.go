package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

func GetUserRBAC(c echo.Context) error {
	user := User{}
	id := c.Param("id")

	db := db.Session()
	if err := db.Where("id = ?", id).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is not exist!")
	}

	rbac, ok := auth.GetRBACByRole(user.Role)
	if !ok {
		return echo.ErrInternalServerError
	}

	resp := struct {
		Success bool     `json:"success"`
		Rbac    []string `json:"rbac"`
	}{
		Success: true,
		Rbac:    rbac,
	}

	return c.JSON(http.StatusOK, &resp)
}
