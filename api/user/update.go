package user

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

// EditUser is used to edit user information by ID
// Allow update values: name, email, username, timezone
// password and role will be update in other API endpoint because of security reason!
func UpdateUser(c echo.Context) error {
	oldUser := User{}
	id := c.Param("id")

	db := db.Session()
	if err := db.Where("id = ?", id).First(&oldUser).Error; err == gorm.ErrRecordNotFound {
		return echo.NewHTTPError(http.StatusBadRequest, "User ID not found!")
	}

	newUser := User{}
	if err := c.Bind(&newUser); err != nil {
		return echo.ErrBadRequest
	}

	// Check if new email address and username is already used by other user!
	if newUser.Email != "" && newUser.Email != oldUser.Email {
		if ok, err := ValidateField(newUser.Email, "email"); !ok {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		checkUser := User{}
		db.Where("email = ?", newUser.Email).First(&checkUser)
		if checkUser.Email == newUser.Email {
			return echo.NewHTTPError(http.StatusBadRequest, "The new email address is already used!")
		}
		if err := db.Model(&newUser).Where("email = ?", oldUser.Email).Update("email", newUser.Email).Error; err != nil {
			return echo.ErrInternalServerError
		}
	}

	if newUser.Username != "" && newUser.Username != oldUser.Username {
		if ok, err := ValidateField(newUser.Username, "username"); !ok {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		checkUser := User{}
		db.Where("username = ?", newUser.Username).First(&checkUser)
		if checkUser.Username == newUser.Username {
			return echo.NewHTTPError(http.StatusBadRequest, "The new username is already used!")
		}
		if err := db.Model(&newUser).Where("username = ?", oldUser.Username).Update("username", newUser.Username).Error; err != nil {
			return echo.ErrInternalServerError
		}
	}

	if newUser.Password != "" {
		if ok, err := ValidateField(newUser.Password, "password"); !ok {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		hashPw, err := HashPassword(newUser.Password)
		if err != nil {
			return echo.ErrInternalServerError
		}
		newUser.PasswordHash = hashPw
		newUser.Password = ""
		if err := db.Model(&newUser).Where("id = ?", id).Update("password_hash", hashPw).Error; err != nil {
			return echo.ErrInternalServerError
		}
	}

	if err := db.Model(&newUser).Where("id = ?", id).Omit("id", "email", "username", "role", "password_hash").Update(&newUser).First(&newUser).Error; err != nil {
		return echo.ErrInternalServerError
	}

	return c.JSON(http.StatusOK, &newUser)
}
