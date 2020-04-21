package user

import (
	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/auth"
	"github.com/nguyentrungtin/go-echo-boilerplate/model"
)

type (
	User          = model.User
	PasswordReset = model.PasswordReset
)

func Router(e *echo.Echo) {
	// ROLE BASE ACCESS CONTROL (RBAC) is setting on the auth package
	// PERMISSION ORDER: SUPER -> ADMIN -> USER

	// DEVELOPER is the special role
	e.POST("/api/developer/create", Developer, auth.Super())

	// ADMIN can manage USER with any role, but not SUPER
	e.GET("/api/user/all", GetAll, auth.Admin())
	e.POST("/api/user/create", CreateUser)

	e.POST("/api/user/login", Login)

	e.GET("/api/user/rbac/all", GetAllRBAC)
	e.GET("/api/user/rbac/:id", GetUserRBAC)

	e.PUT("/api/user/password", ChangePassword)
	e.POST("/api/user/forgot_password", ForgotPassword)
	e.POST("/api/user/reset_password", ResetPassword)

	// ADMIN can udpate role, toggle active and delete user
	e.PUT("/api/user/role/:id", UpdateUserRole, auth.Admin())
	e.PUT("/api/user/toggle_active/:id", ToggleUserActive, auth.Admin())
	e.DELETE("/api/user/:id", DeleteUser, auth.Admin())

	e.GET("/api/user/:id", GetUser)
	e.PUT("/api/user/:id", UpdateUser)
}
