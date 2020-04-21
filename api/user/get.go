package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

// GetUser is used to get one user detail by ID
func GetUser(c echo.Context) error {
	user := User{}
	id := c.Param("id")

	db := db.Session()
	if err := db.Where("id = ?", id).Preload("Dealerships").First(&user).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID is not exist!")
	}

	return c.JSON(http.StatusOK, &user)
}
