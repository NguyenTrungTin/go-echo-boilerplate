package user

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

// CreateUser is used to create a new user
func CreateUser(c echo.Context) error {
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
	user.Role = strings.ToUpper(user.Role)
	user.Status = "ACTIVE"

	if err := db.Create(&user).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusCreated, user)
}
