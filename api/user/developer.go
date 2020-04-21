package user

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

func Developer(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return echo.ErrBadRequest
	}

	if ok, err := Validate(&user); !ok {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	db := db.Session()
	u := User{}
	exist := db.Where("email = ?", user.Email).Or("username = ?", user.Username).First(&u)
	if exist.Error == nil && u.Email == user.Email {
		return echo.NewHTTPError(http.StatusBadRequest, "Email is already exist!")
	}
	if exist.Error == nil && u.Username == user.Username {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is already exist!")
	}

	hashPw, err := HashPassword(user.Password)
	if err != nil {
		return echo.ErrInternalServerError
	}
	user.PasswordHash = hashPw
	user.Password = ""
	user.Role = "DEVELOPER"

	db.Create(&user)

	token, err := auth.DeveloperToken(int(user.ID), user.Username, user.Role)
	if err != nil {
		return echo.ErrInternalServerError
	}

	resp := struct {
		Success bool   `json:"success"`
		User    User   `json:"user"`
		Token   string `json:"token"`
	}{
		Success: true,
		User:    user,
		Token:   token,
	}

	return c.JSON(http.StatusCreated, &resp)
}
