package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
	"golang.org/x/crypto/bcrypt"
)

func ChangePassword(c echo.Context) error {
	user := User{}
	update := struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
	}{}

	id, ok := auth.GetUserID(c)
	if !ok {
		return echo.ErrBadRequest
	}

	if err := c.Bind(&update); err != nil {
		return echo.ErrBadRequest
	}

	db := db.Session()
	result := db.Where("id = ?", id).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not found!")
	}

	if ok, err := ValidateField(update.NewPassword, "password"); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(update.CurrentPassword)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Current password is not correct!")
	}

	hashPw, err := HashPassword(update.NewPassword)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if err := db.Model(&user).Where("id = ?", id).Update("password_hash", hashPw).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return lib.JSON(http.StatusOK, "Password updated successfully!")
}
