package status

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nguyentrungtin/go-echo-boilerplate/db"
)

type Status struct {
	App        string    `json:"app"`
	Database   string    `json:"database"`
	Timestamps time.Time `json:"timestamps"`
}

// Healthz is used to check health status of the application
func Healthz(c echo.Context) error {
	db := db.Session()
	if err := db.DB().Ping(); err != nil {
		s := &Status{
			App:        "FALURE",
			Database:   "FAILURE",
			Timestamps: time.Now().UTC(),
		}
		return c.JSON(http.StatusInternalServerError, s)
	}
	s := &Status{
		App:        "SUCCESS",
		Database:   "SUCCESS",
		Timestamps: time.Now().UTC(),
	}
	return c.JSON(http.StatusOK, s)
}
