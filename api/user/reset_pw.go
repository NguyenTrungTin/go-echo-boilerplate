package user

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

var Reset_Password_Expired int64 = 64

func ResetPassword(c echo.Context) error {
	user := User{}
	password_reset := PasswordReset{}

	data := struct {
		NewPassword string `json:"new_password"`
		Token       string `json:"token"`
	}{}

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	db := db.Session()
	if err := db.Where("token = ?", data.Token).First(&password_reset).Error; gorm.IsRecordNotFoundError(err) {
		return echo.NewHTTPError(http.StatusBadRequest, "Token is not valid!")
	}

	deadline := password_reset.CreatedAt.Add(time.Duration(Reset_Password_Expired) * time.Hour)
	isExpired := time.Now().After(deadline)

	if password_reset.Token == "" || isExpired {
		return echo.NewHTTPError(http.StatusBadRequest, "Reset password token is invalid or expired!")
	}

	// Clean token and soft delete the record
	if err := db.Model(&password_reset).Update("token", "").Delete(&password_reset).Error; err != nil {
		return echo.ErrInternalServerError
	}

	// Update new password for user
	hashPw, err := HashPassword(data.NewPassword)
	if err != nil {
		return echo.ErrInternalServerError
	}
	user.PasswordHash = hashPw

	if err := db.Model(&user).Where("email = ?", password_reset.Email).Update("password_hash", hashPw).First(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	token, err := auth.NewToken(int(user.ID), user.Username, user.Role)
	if err != nil {
		return echo.ErrInternalServerError
	}

	resp := struct {
		Success bool   `json:"success"`
		Message string `json:"messsage"`
		User    User   `json:"user"`
		Token   string `json:"token"`
	}{
		Success: true,
		Message: "Reset password successfully!",
		User:    user,
		Token:   token,
	}
	return c.JSON(http.StatusOK, &resp)
}
