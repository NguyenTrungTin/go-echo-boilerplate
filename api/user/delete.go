package user

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
)

// DeleteUser is used to delete user by ID
func DeleteUser(c echo.Context) error {
	user := User{}
	id := c.Param("id")

	db := db.Session()
	if check := db.Where("id = ?", id).First(&user).Error; check == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not found!")
	}

	// Soft delete, record is still exist in database. The value deleted_at will be set
	if err := db.Where("id = ?", id).Delete(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return lib.JSON(http.StatusOK, fmt.Sprintf("User with ID %s has been delete successfully!", id))
}
