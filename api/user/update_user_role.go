package user

import (
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

func UpdateUserRole(c echo.Context) error {
	user := User{}
	id := c.Param("id")

	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}

	db := db.Session()
	if err := db.Where("id = ?", id).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is not exist!")
	}

	user.Role = strings.ToUpper(user.Role)
	if err := db.Model(&user).Where("id = ?", id).Update("role", user.Role).First(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, &user)
}
