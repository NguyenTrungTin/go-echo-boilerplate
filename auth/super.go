package auth

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
	"github.com/nguyentrungtin/go-echo-boilerplate/model"
)

func Super() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id, ok := GetUserID(c)
			if !ok {
				return echo.ErrUnauthorized
			}

			user := model.User{}
			db := db.Session()
			if err := db.Where("id = ?", id).First(&user).Error; gorm.IsRecordNotFoundError(err) {
				return echo.NewHTTPError(http.StatusBadGateway, "User is no longer exist!")
			}

			if _, ok := lib.Find(RBAC.SuperRBAC, user.Role); !ok {
				return echo.ErrForbidden
			}

			return next(c)
		}
	}
}
