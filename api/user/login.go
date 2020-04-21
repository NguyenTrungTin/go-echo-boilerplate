package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"golang.org/x/crypto/bcrypt"
)

// Login is used to login user in
// If email and password are correct, response token and user role
// Otherwise, response bad request
func Login(c echo.Context) error {
	user := User{}
	newUser := User{}

	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	db := db.Session()
	check := db.Where("email = ?", newUser.Login).Or("username = ?", newUser.Login).First(&user)
	if check.Error == gorm.ErrRecordNotFound || check.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid credentials!")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(newUser.Password)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid credentials!")
	}

	if user.Status == "INACTIVE" {
		return echo.NewHTTPError(http.StatusBadRequest, "Sorry your account is inactive and may not login!")
	}

	var token string
	var err error
	if user.Role == "DEVELOPER" {
		token, err = auth.DeveloperToken(int(user.ID), user.Username, user.Role)
	} else {
		token, err = auth.NewToken(int(user.ID), user.Username, user.Role)
	}
	if err != nil {
		return echo.ErrInternalServerError
	}

	resp := struct {
		Success bool   `json:"success"`
		Token   string `json:"token"`
		Role    string `json:"role"`
	}{
		Success: true,
		Token:   token,
		Role:    user.Role,
	}

	return c.JSON(http.StatusOK, &resp)
}
