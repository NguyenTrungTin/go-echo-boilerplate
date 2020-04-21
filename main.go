package main

import (
	"fmt"

	"github.com/nguyentrungtin/go-echo-boilerplate/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nguyentrungtin/go-echo-boilerplate/api/status"
	"github.com/nguyentrungtin/go-echo-boilerplate/api/user"
	"github.com/nguyentrungtin/go-echo-boilerplate/config"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
	"github.com/nguyentrungtin/go-echo-boilerplate/lib"
)

func main() {
	db.Init()

	if config.APP_ENV == "local" {
		db.LogMode(true)
	}

	//db.Migrate()
	db.AutoMigrate()
	db.Relation()
	defer db.Close()

	e := echo.New()
	e.Debug = true

	e.Use(middleware.CORS())
	e.Use(lib.Logger())
	e.Use(lib.Recover())

	e.Use(auth.JWT())

	user.Router(e)
	status.Router(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.PORT)))
}
