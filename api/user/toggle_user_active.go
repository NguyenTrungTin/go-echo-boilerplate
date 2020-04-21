package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

func ToggleUserActive(c echo.Context) error {
	user := User{}
	id := c.Param("id")

	db := db.Session()
	if err := db.Model(&user).Where("id = ?", id).First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is not exist!")
	}

	if user.Status == "ACTIVE" {
		user.Status = "INACTIVE"
	} else if user.Status == "INACTIVE" {
		user.Status = "ACTIVE"
	} else {
		return echo.ErrBadRequest
	}

	if err := db.Save(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, &user)
}
