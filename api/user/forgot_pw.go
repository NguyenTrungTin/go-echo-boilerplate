package user

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/api/email"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
	"github.com/nguyentrungtin/go-echo-boilerplate/model"
)

func ForgotPassword(c echo.Context) error {
	user := User{}
	data := struct {
		Email string `json:"email"`
	}{}

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	if ok, err := ValidateField(data.Email, "email"); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	db := db.Session()
	check := db.Where("email = ?", data.Email).First(&user)
	if check.Error != nil && check.Error != gorm.ErrRecordNotFound {
		return echo.ErrInternalServerError
	}

	// Delete all tokens that belong to this email if exist to prevent duplicate!
	if err := db.Model(PasswordReset{}).Where("email = ?", data.Email).Updates(map[string]interface{}{"token": ""}).Delete(PasswordReset{}).Error; err != nil {
		return echo.ErrInternalServerError
	}

	if user.Email == data.Email {
		b := make([]byte, 64)
		if _, err := rand.Read(b); err != nil {
			return echo.ErrInternalServerError
		}

		token := fmt.Sprintf("%x", b)

		pw_reset := model.PasswordReset{
			UserID:    user.ID,
			Email:     user.Email,
			Token:     token,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}

		go email.SendResetPassword(user.Email, user.FirstName, token)

		if err := db.Create(&pw_reset).Error; err != nil {
			return echo.ErrInternalServerError
		}
	}

	return lib.JSON(http.StatusOK, "We'll sent the reset password link to your inbox if your email exist in our system!")
}
